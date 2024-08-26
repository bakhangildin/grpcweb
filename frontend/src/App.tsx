import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { Component } from "solid-js";
import { MyHelloServiceClient } from "./contracts/api.client";
import { HelloRequest } from "./contracts/api";

const App: Component = () => {
  const transport = new GrpcWebFetchTransport({
    format: "binary",
    baseUrl: "http://localhost:4000/grpc/",
    fetchInit: {
      credentials: "include",
    },
  });
  const cl = new MyHelloServiceClient(transport);
  async function handleSubmit(e: SubmitEvent) {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const result = form.querySelector<HTMLPreElement>("#result");
    if (!result) return;
    try {
      const myName = form.myName.value;
      const myAge = parseInt(form.myAge.value);
      const r = await cl.hello(
        HelloRequest.create({
          myName: myName,
          myAge: myAge,
        }),
      );
      result.innerHTML = JSON.stringify(r.response, null, 2);
    } catch (e) {
      result.innerHTML = `${e}`;
    }
  }

  return (
    <form onSubmit={handleSubmit} autocomplete="off">
      <input
        class="border block"
        type="text"
        name="myName"
        placeholder="name"
        required
      />
      <input
        class="border block"
        type="text"
        name="myAge"
        placeholder="age"
        required
      />
      <button type="submit">Call MyHelloService.Hello</button>
      <pre id="result"></pre>
    </form>
  );
};

export default App;
