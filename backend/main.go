package main

import (
	"backend/contracts"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
)

type MyHelloService struct {
	contracts.UnimplementedMyHelloServiceServer
}

func (s *MyHelloService) Hello(ctx context.Context, in *contracts.HelloRequest) (*contracts.HelloResponse, error) {
	now := time.Now().UTC()
	out := &contracts.HelloResponse{
		ResponseText:      fmt.Sprintf("Hello from server '%s' your age is %d", in.GetMyName(), in.GetMyAge()),
		ServerUnixTimeI:   int32(now.Unix()),
		ServerUnixTimeStr: now.Format(time.DateTime),
	}
	return out, nil
}

func (s *MyHelloService) HelloStream(in *contracts.HelloRequest, stream grpc.ServerStreamingServer[contracts.HelloResponse]) error {
	i := 0
	n := 10
	t := time.NewTicker(time.Second)
	for {
		select {
		case <-t.C:
			i++
			if i >= n {
				fmt.Println("no iters left")
				return nil
			}
			if i == 4 {
				return fmt.Errorf("i == %d", i)
			}
			now := time.Now().UTC()
			if err := stream.Send(&contracts.HelloResponse{
				ResponseText:      fmt.Sprintf("Hello from server '%s' your age is %d", in.GetMyName(), in.GetMyAge()),
				ServerUnixTimeI:   int32(now.Unix()),
				ServerUnixTimeStr: now.Format(time.DateTime),
			}); err != nil {
				fmt.Println("send error", err)
				return err
			}
		case <-stream.Context().Done():
			fmt.Println("context cancelled")
			return nil
		}
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Add("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
		next.ServeHTTP(w, r)
	})

}

type GRPCHandler struct {
	grpcWebServer *grpcweb.WrappedGrpcServer
}

func (h *GRPCHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == http.MethodPost {
		h.grpcWebServer.ServeHTTP(w, r)
		return
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("%s only for grpc-web calls", r.URL.Path)))
		return
	}
}

func main() {
	s := grpc.NewServer()
	grpcWebServer := grpcweb.WrapServer(s)
	contracts.RegisterMyHelloServiceServer(s, &MyHelloService{})
	mux := chi.NewMux()
	mux.Use(
		middleware.Logger,
		corsMiddleware,
	)
	h := &GRPCHandler{
		grpcWebServer: grpcWebServer,
	}
	mux.Handle("/*", http.FileServerFS(os.DirFS("../frontend/dist/")))
	mux.Handle("/grpc/*", http.StripPrefix("/grpc", h))
	log.Fatal(http.ListenAndServe(":4000", mux))
}
