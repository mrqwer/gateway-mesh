package main

import (
	"context"
	"log"
	"net"

	pb "github.com/mrqwer/gateway-mesh/proto/users"
	"google.golang.org/grpc"
)

type usersServer struct {
	pb.UnimplementedUsersServiceServer
}

func (s *usersServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{
		Id:    req.Id,
		Name:  "John Doe",
		Email: "john@example.com",
	}, nil
}

func (s *usersServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users := []*pb.GetUserResponse{
		{Id: "1", Name: "John", Email: "john@example.com"},
		{Id: "2", Name: "Jane", Email: "jane@example.com"},
	}
	return &pb.ListUsersResponse{Users: users}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUsersServiceServer(s, &usersServer{})

	log.Println("users service listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
