package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	pb "server/proto/user"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.User, error) {
	// Here, we would normally fetch the user from a database or other data source
	// based on the ID in the request, but for simplicity, we'll just return a
	// hard-coded user.
	user := &pb.User{
		Id:   req.Id,
		Name: "John Doe",
	}
	return user, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
