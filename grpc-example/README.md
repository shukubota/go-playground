# example
```shell
grpcurl -plaintext localhost:50051 list
```

でmethod一覧取得。

```sh
grpcurl -plaintext -d '{"name": "test"}' localhost:50051 hello.Greeter.SayHello
```

## 準備
protoc install (brewで良い)
```shell
protoc --version
// libprotoc 3.21.5
```

## コンパイル
protoファイルをコンパイルする
```shell
make protoc
```

## curl
```shell
grpcurl -plaintext -d '{"name": "hoge"}' localhost:50051 hello.Greeter.SayHello
```
で正常に帰ってくる