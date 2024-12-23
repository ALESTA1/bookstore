package main

import (
	pb "bookstore/proto"
	"bookstore/server/auth"
	"bookstore/server/interceptor"
	"context"
	"log"
	"net"
	"os"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.BookServiceServer
	AuthSvc *auth.Service
}

func UnaryServerInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	log.Printf("Received request on method: %s", info.FullMethod)
	resp, err := handler(ctx, req)
	log.Printf("Sending response from method: %s", info.FullMethod)
	return resp, err
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatalf("SECRET_KEY environment variable is not set")
	}

	authSvc, err := auth.NewService(secretKey) // will be stored as env variable in prod
	if err != nil {
		log.Fatalf("failed to initialize auth service: %v", err)
	}
	interceptor, err := interceptor.NewAuthInterceptor(authSvc)
	if err != nil {
		log.Fatalf("failed to initialize interceptor: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.UnaryAuthMiddleware),
	)

	pb.RegisterBookServiceServer(grpcServer, &helloServer{AuthSvc: authSvc})

	reflection.Register(grpcServer) //for testing with grpcurl

	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
