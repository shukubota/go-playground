package main

import (
	"fmt"
	"github.com/shukubota/go-api-template/serverless-template/searchContent/handler"
	"log"
)

func main() {
	r, err := handler.Handler()
	fmt.Println(r)
	if err != nil {
		log.Fatalln("handler error")
	}
}
