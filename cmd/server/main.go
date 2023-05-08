package main

import (
	"fmt"
	"log"
	"net"
	"os"

	pb "test/gen/api/v1"
	"test/pkg/grpc/server/controllers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	host = "localhost"
	port = "50051"
)

func main() {
	
	addr := fmt.Sprintf("%s:%s", host, port)
	lis, err := net.Listen("tcp", addr)	
	if err != nil {
		log.Println("error starting tcp listener: ", err)
		os.Exit(1)
	}

	log.Println("tcp listener started at port: ", port)
	grpcServer := grpc.NewServer()
	ctServer := controllers.NewServer()
	pb.RegisterBasicServiceServer(grpcServer, ctServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Println("error serving grpc: ", err)
		os.Exit(1)
	}

}