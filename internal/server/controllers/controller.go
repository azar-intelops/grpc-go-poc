package controllers

import (
	"context"
	"test/gen/pb"
)

type Server struct {
	pb.UnimplementedBasicServiceServer
}

func NewServer() *Server  {
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