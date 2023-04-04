package main

import (
	"context"
	"fmt"

	pb "client/proto/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	user, err := client.GetUser(context.Background(), &pb.UserRequest{Id: 123})
	if err != nil {
		panic(err)
	}

	fmt.Printf("User: %+v\n", user)
}
