package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	doSomething(ctx, nil)
	fmt.Println("goroutines:", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	fmt.Println("goroutines:", runtime.NumGoroutine())
}

func doSomething(ctx context.Context, ss <-chan string) <-chan any {
	completed := make(chan any)

	go func() {
		defer fmt.Println("doSomething exited")
		defer close(completed)
		for {
			select {
			case s := <-ss:
				fmt.Println(s)
			// case <-time.After(1 * time.Second):
			// 	return
			case <-ctx.Done():
				return
			}
		}
	}()

	return completed
}
