# UPDATE

Starting from [version 1.26](https://github.com/serverless/serverless/releases/tag/v1.26.0) Serverless Framework includes two Golang templates:

* `aws-go` - basic template with two functions
* `aws-go-dep` - **recommended** template using [`dep`](https://github.com/golang/dep) package manager

You can use them with `create` command:

```
serverless create -t aws-go-dep
```

Original README below.

---

# Serverless Template for Golang

This repository contains template for creating serverless services written in Golang.

## Quick Start

1. Create a new service based on this template

```
serverless create -u https://github.com/serverless/serverless-golang/ -p serverless-template
```

2. Compile function

```shell
cd serverless-template
GOOS=linux go build -o bin/main

# scraping用
GOOS=linux go build -o bin/searchContent ./searchContent

# deviceToken用
GOOS=linux go build -o bin/registerDeviceToken ./registerDeviceToken

# line用
GOOS=linux go build -o bin/line ./line
```

3. Deploy!

```
serverless deploy --region ap-northeast-1
```
