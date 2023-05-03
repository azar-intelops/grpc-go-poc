package controllers

import (
	"context"

	"github.com/azar-intelops/grpc-code-gen/server/pb"
)

type Server struct {
	pb.UnimplementedBasicServiceServer
	// pb.Unimplemented{{.ServiceName}}Server
}

func NewServer()  *Server {
	return &Server{}
}

func (s *Server) Get(ctx context.Context,req *pb.MessageRequest) (*pb.MessageResponse, error)  {
	return &pb.MessageResponse{
		Result: "Get Success",
	}, nil
}

func (s *Server) Create(ctx context.Context,req *pb.MessageRequest) (*pb.MessageResponse, error)  {
	return &pb.MessageResponse{
		Result: "Create Success",
	}, nil
}

func (s *Server) Delete(ctx context.Context,req *pb.MessageRequest) (*pb.MessageResponse, error)  {
	return &pb.MessageResponse{
		Result: "Delete Success",
	}, nil
}

func (s *Server) Update(ctx context.Context,req *pb.MessageRequest) (*pb.MessageResponse, error)  {
	return &pb.MessageResponse{
		Result: "Update Success",
	}, nil
}