package main

import (
	"context"
	service "grpc-spanner-app/gen/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	service.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *service.GetUserRequest) (*service.GetUserResponse, error) {
	return &service.GetUserResponse{
		UserId: req.GetUserId(),
		Name:   "John Doe",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
