package main

import (
	"go-grpc/cmd/config"
	"go-grpc/cmd/services"
	productPb "go-grpc/pb/product"
	"log"
	"net"

	"github.com/lpernett/godotenv"
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

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Cant load env file")
		return
	}

	db := config.Get()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Server started at %v\n", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v\n", err.Error())
	}
}
