package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func UnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Println("=============unary interceptor")
	return handler(ctx, req)
}
