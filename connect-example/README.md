## 初期セットアップ
https://connect.build/docs/go/getting-started

```shell
$ go get github.com/bufbuild/buf/cmd/buf@latest
$ go get github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
$ go get google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go get github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

## リクエスト
```shell
curl -H "Content-Type: application/json" -d '{"name": "Jane"}' http://localhost:18080/greet.v1.GreetService/Greet
```
reqponse↓
```shell
{"greeting":"Hello Jane"}
```