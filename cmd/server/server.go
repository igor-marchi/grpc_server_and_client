package main

import (
	"fmt"
	"log"
	"net"

	"github.com/igor-marchi/grpc_full_cycle/pb"
	"github.com/igor-marchi/grpc_full_cycle/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// Run server
// go run cmd/server/server.go

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	reflection.Register(grpcServer)

	fmt.Println("Server run 🔥")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %v", err)
	}
}
