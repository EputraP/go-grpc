package main

import (
	"context"
	"fmt"
	productPb "go-grpc/pb/product"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("start connection")

	netListen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen %v\n", err.Error())
	}

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error")
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := productPb.NewProductServiceClient(conn)

	status, err := client.DeleteProduct(context.Background(), &productPb.Id{Id: 4})
	if err != nil {
		fmt.Println("error")
		log.Fatalf("did not connect: %v", err)
	}

	fmt.Println(status.Status)

	grpcServer := grpc.NewServer()

	log.Printf("Server started at %v\n", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve %v\n", err.Error())
	}
}
