package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func RecoverUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("=============recover interceptor")
	fmt.Println(info)
	return handler(ctx, req)
}
