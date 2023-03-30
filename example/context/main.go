package main

import (
	"context"
	"fmt"
	"time"
)

func parentFunc(ctx context.Context) {
	fmt.Println("start parentFunc")
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("parentFunc canceled")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
	// cancelされると到達しない
	fmt.Println("parentFunc finished")
}

func childFunc(ctx context.Context) {
	fmt.Println("start childFunc")
	for i := 1; i <= 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("childFunc canceled")
			return
		default:
			time.Sleep(1 * time.Second)
		}
	}
	// parentのcancelはここにも伝搬するので到達しない
	fmt.Println("childFunc finished")
}

func main() {
	fmt.Println("start main")
	defer fmt.Println("done main")
	ctx := context.Background()

	parent, cancel := context.WithCancel(ctx)
	go parentFunc(parent)

	child, _ := context.WithCancel(parent)
	go childFunc(child)

	time.Sleep(1 * time.Second)

	cancel()

	time.Sleep(3 * time.Second)
}
