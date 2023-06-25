import { createPromiseClient } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

import { GreetService } from "../gen/greet/v1/greet_connect";

import './App.css'
import {ServiceType} from "@bufbuild/protobuf";
import useSWR from "swr";

const transport = createConnectTransport({
  baseUrl: "http://localhost:18888",
});

const client = createPromiseClient(GreetService as ServiceType, transport);

const fetchData = async () => {
  const response = await fetch('http://localhost:8085/example-messages');
  response.json().then((data) => {
    console.log(data);
    return data;
  }).catch((err) => {
    console.log(err);
  });
}

function App() {
  const onFeth = async () => {
    const res = await fetchData();
    console.log({ res })
  }

  const { data, error, isLoading } = useSWR("greet", () => {
    return client.greet({ name: "buf" });
    // return client.greet({ name: "buf" });
  });

  console.log({ data, error, isLoading })

  return (
    <>
      <div>
        {isLoading && <p>loading...</p>}
        {error && <p>error: {error.message}</p>}
        {data && <p>data: {data.greeting}</p>}
      </div>
      <button onClick={onFeth}>fetch data</button>
    </>
  )
}

export default App
