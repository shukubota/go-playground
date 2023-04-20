package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 3 * time.Second

func process1(ctx context.Context) {
	time.Sleep(2 * time.Second)

}

func main() {
	ctx2 := context.TODO()
	fmt.Println(ctx2.Value("aaa"))
	ctx3 := context.WithValue(ctx2, "aaa", struct{ name string }{name: "aaa"})

	fmt.Println(ctx3.Value("aaa"))

	//d := time.Now().Add(shortDuration)
	//ctx, cancel := context.WithDeadline(context.Background(), d)

	ctx, cancel := context.WithTimeout(context.Background(), shortDuration)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	select {
	case <-time.After(4 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(context.Cause(ctx))
	}

}
