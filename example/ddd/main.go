package main

import (
	"github.com/shukubota/go-api-template/example/ddd/repository"
	"github.com/shukubota/go-api-template/example/ddd/usecase"
	"log"
)

func main() {
	name := "Bob"
	uc := usecase.NewRegisterUserUseCase(repository.NewRegisterUserRepository())

	err := uc.Register(name)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
}
