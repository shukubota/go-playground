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
curl -H "Content-Type: application/json" -H "Origin: localhost:5173" -d '{"name": "Jane"}' http://localhost:18888/greet.v1.GreetService/Greet
```
reqponse↓
```shell
{"greeting":"Hello Jane"}
```

grpcurlの場合↓でいける
```shell
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"name": "Jane"}' \
    localhost:18080 greet.v1.GreetService/Greet
```

```shell
buf generate --template buf.gen.yaml (or buf.gen.js.yaml)
```

## bufコンパイル
```shell
buf generate
```
でgo側のgenディレクトリ、frontend/genディレクトリにファイルが生成される

## サーバ起動
```shell
go run main.go
```
でconnect用のサーバがたつ

```shell
npm run dev
```
でフロントがたち、そこからlocalhost:18888へconnect-webのリクエストがくる

fetch dataは別途grpc-playgroundのサーバを立てる必要がある

## リクエスト
### 素のgrpcに対するgrpcurl
```shell
grpcurl -plaintext -d '{"name": "hoge"}' localhost:29999 greet.v1.GreetService.Greet
```
