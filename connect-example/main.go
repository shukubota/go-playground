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
	//h2 := h2c.NewHandler(mux, &http2.Server{})

	fmt.Println(h)

	port := 18888
	log.Printf("Starting server on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), h))
}
