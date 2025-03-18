package healthv1

import (
	"context"
	"log"

	pb "github.com/shukubota/go-playground/grpc-gateway-example/api/health/v1"
)

// HealthServer はヘルスチェックサービスの実装
type HealthServer struct {
	pb.UnimplementedHealthServiceServer
}

// NewHealthServer は新しいHealthServerを作成します
func NewHealthServer() *HealthServer {
	return &HealthServer{}
}

// CheckHealth はヘルスチェックリクエストを処理します
func (s *HealthServer) CheckHealth(ctx context.Context, req *pb.CheckHealthRequest) (*pb.CheckHealthResponse, error) {
	memberID := ""
	if req.Parameters != nil && req.Parameters.Member != nil {
		memberID = req.Parameters.Member.MemberId
	}

	log.Printf("Received health check request with member_id: %s", memberID)

	return &pb.CheckHealthResponse{
		MemberId: memberID,
	}, nil
}
