package main

import (
	"context"
	"fmt"
	"time"
)

func subroutine(ctx context.Context) {
	fmt.Println("subroutine start...")
	select {
	case <-ctx.Done():
		fmt.Println("subroutine end")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go subroutine(ctx)

	time.Sleep(time.Second * 1)
	cancel()
	time.Sleep(time.Second * 1)
	fmt.Printf("cause err: %+v\n", context.Cause(ctx).Error()) // context canceled (context.Canceled)
}
