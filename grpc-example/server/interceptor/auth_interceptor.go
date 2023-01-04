package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func AuthUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("=============unary interceptor")
	fmt.Println(info)
	return handler(ctx, req)
}
