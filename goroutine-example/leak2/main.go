package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	repo := &someRepositoryImpl{}
	memberChan, errChan := repo.GetMember(ctx)

	select {
	case member := <-memberChan:
		fmt.Println(member)
	case err := <-errChan:
		fmt.Println("error:", err)
		return err
	case <-ctx.Done():
		fmt.Println("context canceled")

		time.Sleep(2 * time.Second)

		// GetMemberの中でctx.Doneを見ないと、goroutineが終了しないまま以下は2になる
		fmt.Println("タイムアウト後2秒後のgoroutine数:", runtime.NumGoroutine())

		return ctx.Err()
	}

	return nil
}

type Member struct {
	ID string `json:"id"`
}

type someRepositoryImpl struct{}

func (s *someRepositoryImpl) GetMember(ctx context.Context) (<-chan *Member, chan error) {
	memberChan := make(chan *Member)
	errChan := make(chan error)
	go func() {
		// defer close(memberChan)
		// defer close(errChan)

		var member *Member
		err := json.Unmarshal([]byte(`a{"id":"123"}`), &member)
		if err != nil {
			errChan <- err
			return
		}
		// time.Sleep(2 * time.Second) // Simulate a long operation

		select {
		case <-ctx.Done():
			return
		case memberChan <- member:
		}
	}()
	return memberChan, errChan
}
