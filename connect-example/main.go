package main

import (
	"connect-example/gen/greet/v1/greetv1connect"
	"connect-example/handler"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"net/http"
)

func main() {
	greetHandler := handler.NewGreetHandler()

	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greetHandler)

	mux.Handle(path, handler)
	http.ListenAndServe("localhost:18080", h2c.NewHandler(mux, &http2.Server{}))
}
