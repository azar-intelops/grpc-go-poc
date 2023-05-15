package main

import (
	"context"
	"log"
	pb "test/gen/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = "localhost"
	port = ":50051"
)

func main() {
	conn, err := grpc.Dial(host+port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
	}

	userClient := pb.NewBasicServiceClient(conn)

	ping, _ := userClient.Ping(context.Background(), &pb.Request{})

	log.Println(ping)

}