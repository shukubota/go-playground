package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	e := status.Error(codes.DeadlineExceeded, "aaa")
	s, ok := status.FromError(e)

	fmt.Println(s)
	fmt.Println(ok)
	fmt.Println(s.Err())

	e = errors.New("test")
	//eWithStack := errors.WithStack
	s, ok = status.FromError(e)

	fmt.Println(s)
	fmt.Println(ok)
	fmt.Println(s.Err())

	e = context.Canceled

	if !errors.Is(e, context.Canceled) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
