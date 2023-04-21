package main

import (
	"context"
	"fmt"
	"time"
)

func subroutine2(ctx context.Context) {
	fmt.Println("subroutine start...")
	select {
	case <-ctx.Done():
		fmt.Println("subroutine end")
		return
	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()

	go subroutine2(ctx)

	time.Sleep(time.Second * 3)
	fmt.Printf("cause err: %+v\n", context.Cause(ctx).Error()) // context deadline exceeded (context.DeadlineExceeded)
}
