package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/m-mizutani/gollem"
	"github.com/m-mizutani/gollem/llm/gemini"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}

func run() error {
	ctx := context.Background()

	projectID := "gig-sandbox-ai"
	model := "gemini-1.5-flash"
	location := "us-central1"

	client, err := gemini.New(ctx, projectID, location, gemini.WithModel(model))
	if err != nil {
		return fmt.Errorf("failed to create Gemini client: %w", err)
	}

	fmt.Println(client)

	ssn, err := client.NewSession(ctx)
	if err != nil {
		return err
	}

	prompt := "What is the capital of Japan?"

	result, err := ssn.GenerateContent(ctx, gollem.Text(prompt))
	if err != nil {
		return err
	}
	fmt.Println("Response:", result.Texts)
	return nil
}
