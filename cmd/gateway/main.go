package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	// "github.com/mrqwer/gateway-mesh/internal/gateway"
	"github.com/mrqwer/gateway-mesh/internal/gateway"
	pb "github.com/mrqwer/gateway-mesh/proto/users"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//  redis

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()

	// gRPC connection options — using insecure for dev only
	connOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// ✅ Match exactly the users service address (as in Docker compose or localhost)
	usersSvcAddr := "localhost:50051" // or "users:50051" if using Docker compose

	if err := pb.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, usersSvcAddr, connOpts); err != nil {
		log.Fatalf("failed to register gateway handler: %v", err)
	}

	// ✅ Attach middlewares (JWT + RateLimiter)
	// handler := gateway.JWTMiddleware(gateway.RateLimitMiddleware(mux))
	handler := gateway.RateLimitMiddleware(mux)

	// handler := mux

	server := &http.Server{
		Addr:         ":8080",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("🚀 API Gateway running on :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("gateway server failed: %v", err)
	}
}
