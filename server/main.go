package main

import (
	"log"
	"net"

	pb "github.com/itsadiyap/Go-GRPC-Server/proto"
	"google.golang.org/grpc"
)

type server struct{
	pb.GreetServiceServer
}

const (
	port = "8080"
)

func main() {
	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})
	log.Printf("Server listening on port %v", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}