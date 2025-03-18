package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strings"

	pb "github.com/shukubota/go-playground/grpc-gateway-example/api/health/v1"
	healthv1 "github.com/shukubota/go-playground/grpc-gateway-example/pkg/health/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	grpcPort = ":8081"
	httpPort = ":8082"
)

func main() {
	// gRPCサーバーの起動
	go runGRPCServer()

	// HTTPゲートウェイの起動
	runHTTPServer()
}

func runGRPCServer() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHealthServiceServer(s, healthv1.NewHealthServer())

	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// カスタムヘッダーマッチャー
func customHeaderMatcher(key string) (string, bool) {
	switch key {
	case "Content-Type":
		return "Content-Type", true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}

// フォームデータをJSONに変換するミドルウェア
func FormToJSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Content-Typeがapplication/x-www-form-urlencodedかチェック
		if r.Method == "POST" && strings.HasPrefix(r.Header.Get("Content-Type"), "application/x-www-form-urlencoded") {
			// フォームデータをパース
			if err := r.ParseForm(); err != nil {
				http.Error(w, "Failed to parse form data", http.StatusBadRequest)
				return
			}

			// パラメータを抽出
			memberID := ""
			if memberJson := r.FormValue("parameters[member]"); memberJson != "" {
				var member map[string]interface{}
				if err := json.Unmarshal([]byte(memberJson), &member); err == nil {
					if id, ok := member["member_id"].(string); ok {
						memberID = id
					}
				}
			}

			// 新しいリクエストボディを作成
			newBody := map[string]interface{}{
				"parameters": map[string]interface{}{
					"member": map[string]interface{}{
						"member_id": memberID,
					},
				},
			}

			// JSONにエンコード
			jsonBody, err := json.Marshal(newBody)
			if err != nil {
				http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
				return
			}

			// 新しいリクエストを作成
			newReq, err := http.NewRequestWithContext(
				r.Context(),
				r.Method,
				r.URL.String(),
				io.NopCloser(strings.NewReader(string(jsonBody))),
			)
			if err != nil {
				http.Error(w, "Failed to create new request", http.StatusInternalServerError)
				return
			}

			// ヘッダーをコピー
			for k, v := range r.Header {
				newReq.Header[k] = v
			}
			newReq.Header.Set("Content-Type", "application/json")
			newReq.ContentLength = int64(len(jsonBody))

			// 新しいリクエストで処理を続行
			next.ServeHTTP(w, newReq)
			return
		}

		// 他のリクエストは通常どおり処理
		next.ServeHTTP(w, r)
	})
}

func runHTTPServer() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// HTTPゲートウェイを作成
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				EmitUnpopulated: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
		runtime.WithIncomingHeaderMatcher(customHeaderMatcher),
	)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// gRPCサーバーに接続
	if err := pb.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, "localhost"+grpcPort, opts); err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// フォームデータをJSONに変換するミドルウェアを適用
	handler := FormToJSONMiddleware(mux)

	// HTTPサーバーを起動
	log.Printf("HTTP server listening at %v", httpPort)
	if err := http.ListenAndServe(httpPort, handler); err != nil {
		log.Fatalf("failed to serve HTTP: %v", err)
	}
}
