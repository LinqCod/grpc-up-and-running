package main

import (
	"github.com/linqcod/grpc-up-and-running/productInfo/server/ecommerce"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error while listening tcp connection: %v", err)
	}

	s := grpc.NewServer()
	ecommerce.RegisterProductInfoServer(s, &server{})

	log.Printf("Starting gRPC listener on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
