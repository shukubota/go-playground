# go-playground

```sh
docker-compose up
docker-compose exec app bash
```

## go1.20rc1
```shell
go1.20rc1 vet ./sandbox/main_test.go  
```

## install
```shell
$ go install golang.org/dl/go1.20@latest
$ go1.20 download
$ go1.20 version
```

### db
```shell
cd mysql
docker-compose up
mysql -uroot -h 127.0.0.1 -P 13306 -proot_password
```