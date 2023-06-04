import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

import { GreetService } from "../gen/greet/v1/greet_connect";

import './App.css'
import {ServiceType} from "@bufbuild/protobuf";

const transport = createConnectTransport({
  baseUrl: "http://localhost:18080",
});

const client = createPromiseClient(GreetService as ServiceType, transport);

function App() {
  const onSubmit = async () => {
    console.log("onsubmit")
    const res = await client.greet({ name: "buf" });
    console.log({ res });
  }

  return (
    <>
      <p>hello</p>
        <button onClick={onSubmit}>submit</button>
    </>
  )
}

export default App
