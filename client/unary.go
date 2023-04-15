package main

import (
	"context"
	"log"
	"time"

	pb "github.com/itsadiyap/Go-GRPC-Server/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("Error while calling SayHello RPC: %v", err)
	}

	log.Printf("Response from SayHello: %v", res.Message)
}