package main

import (
	"go-grpc/cmd/services"
	productPb "go-grpc/pb/product"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err.Error())
	}

	grpcServer := grpc.NewServer()
	productService := services.ProductService{}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v\n", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v\n", err.Error())
	}
}
