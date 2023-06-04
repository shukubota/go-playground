package main

import (
	"connect-example/gen/greet/v1/greetv1connect"
	"connect-example/handler"
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	greetHandler := handler.NewGreetHandler()

	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greetHandler)

	mux.Handle(path, handler)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173"},
		//AllowOriginFunc: func(origin string) bool {
		//	return true
		//},
		AllowedHeaders: []string{"*"},
	})

	fmt.Println(c)

	h := c.Handler(h2c.NewHandler(mux, &http2.Server{}))

	log.Println("Starting server on port :18080")
	log.Fatal(http.ListenAndServe("localhost:18080", h))
}
