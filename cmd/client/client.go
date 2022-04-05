package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/igor-marchi/grpc_full_cycle/pb"
	"google.golang.org/grpc"
)

// executar client
// go run cmd/client/client.go

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)
	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	// addUser(client)
	// AddUserVerbose(client)
	AddUsers(client)
}

func addUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "João",
		Email: "j@j",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		{
			Id:    "0",
			Name:  "João",
			Email: "j@j",
		},
		{
			Id:    "1",
			Name:  "João1",
			Email: "j1@j",
		},
		{
			Id:    "2",
			Name:  "João2",
			Email: "j2@j",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	fmt.Println(res)
}
