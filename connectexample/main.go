package main

import (
	"connectexample/gen/greet/v1/v1connect"
	"connectexample/handler"
	"context"
	"fmt"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

func main() {
	greetHandler := handler.NewGreetHandler()

	mux := http.NewServeMux()
	path, handler := v1connect.NewGreetServiceHandler(greetHandler)

	mux.Handle(path, handler)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		//w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5174"},
		//AllowOriginFunc: func(origin string) bool {
		//	return true
		//},
		AllowedHeaders: []string{"*"},
	})
	//
	//fmt.Println(c)

	h := c.Handler(h2c.NewHandler(mux, &http2.Server{}))
	//h2 := h2c.NewHandler(mux, &http2.Server{})

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		port := 18888
		log.Printf("Starting server on port %d", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), h))
		return nil
	})

	err := eg.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
