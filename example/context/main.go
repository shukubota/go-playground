package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 4 * time.Second

func process1(ctx context.Context) {
	time.Sleep(2 * time.Second)

}

func main() {
	ctx2 := context.TODO()
	fmt.Println(ctx2.Value("aaa"))
	ctx3 := context.WithValue(ctx2, "aaa", struct{ name string }{name: "aaa"})

	fmt.Println(ctx3.Value("aaa"))

	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)
	defer cancel()

	select {
	case <-time.After(5 * time.Second):
		fmt.Println("overslept")
		fmt.Println(ctx.Err())
	case <-ctx.Done():
		fmt.Println(context.Cause(ctx))
		fmt.Println(ctx.Err())
	}

}
