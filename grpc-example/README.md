# example
```shell
grpcurl -plaintext localhost:50051 list
```

でmethod一覧取得。

```sh
grpcurl -plaintext -d '{"name": "test"}' localhost:50051 hello.Greeter.SayHel
lo
```