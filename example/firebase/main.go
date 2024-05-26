package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		log.Fatalf("error: no uid specified: execute go run main.go <uid>")
	}

	uid := args[0]

	ctx := context.Background()
	e, err := setup(ctx, uid)
	if err != nil {
		log.Fatalf("error: %+v", err)
	}

	log.Printf("email: %s", *e)

	return
}

func setup(ctx context.Context, uid string) (*string, error) {
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		return nil, err
	}

	client, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	u, err := client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}

	return &u.Email, nil
}
