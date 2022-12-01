package main

import (
	"github.com/shukubota/go-api-template/example/ddd/repository"
	"github.com/shukubota/go-api-template/example/ddd/use_case"
	"log"
)

func main() {
	name := "Bob"
	uc := use_case.NewRegisterUserUseCase(repository.NewRegisterUserRepository())

	err := uc.Register(name)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}
}
