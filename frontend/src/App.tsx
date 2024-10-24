import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { Component, For, onMount } from "solid-js";
import { createStore, produce } from "solid-js/store";
import { MyHelloServiceClient } from "./contracts/api.client";
import { HelloRequest } from "./contracts/api";

const App: Component = () => {
  const [store, setStore] = createStore({ items: new Array<string>() });
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
      result.innerHTML = e;
    }
  }

  onMount(async () => {
    const r = cl.helloStream(
      HelloRequest.create({
        myName: "test-name",
        myAge: 12,
      }),
    );
    r.responses.onMessage((r) => {
      console.log("got message", r.serverUnixTimeStr, r.responseText);
      setStore(
        produce((s) => {
          s.items.push(r.responseText);
        }),
      );
    });
    r.responses.onError((e) => {
      console.log("error happened", e.message);
      setStore(
        produce((s) => {
          s.items.push(e.message);
        }),
      );
    });
  });

  return (
    <div class="flex gap-2">
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
        <button class="rounded border-1 border-black border" type="submit">
          Call MyHelloService.Hello
        </button>
        <pre id="result"></pre>
      </form>
      <div>
        <p>Streamed data messages</p>
        <For each={store.items}>{(message) => <div>{message}</div>}</For>
      </div>
    </div>
  );
};

export default App;
