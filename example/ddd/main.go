package main

import (
	"github.com/gorilla/mux"
	"github.com/shukubota/go-api-template/example/ddd/handler"
	"log"
	"net/http"
)

func main() {
	//name := "Bob"
	//uc := usecase.NewRegisterUserUseCase(repository.NewRegisterUserRepository())
	//
	//err := uc.Register(name)
	//if err != nil {
	//	log.Fatalf("error: %+v", err)
	//}

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")

	srv := &http.Server{
		Addr:    "0.0.0.0:5555",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
