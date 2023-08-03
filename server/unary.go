package main

import (
	"context"
	"log"

	pb "github.com/itsadityap/Go-GRPC-Server/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	
	// Log a request is received
	log.Printf("A Request is received from client")

	return &pb.HelloResponse{
		Message: "Hello From Server",
	}, nil
}