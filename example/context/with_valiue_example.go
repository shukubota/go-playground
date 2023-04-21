package main

import (
	"context"
	"fmt"
)

const hogeKey = "hoge"

func subroutine3(ctx context.Context) {
	h := ctx.Value(hogeKey)
	fmt.Println(h)
}

func main() {
	ctx := context.WithValue(context.Background(), hogeKey, "hoge")

	go subroutine3(ctx)
}
