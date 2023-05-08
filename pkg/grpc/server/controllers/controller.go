package controllers

import (
	"context"
	"fmt"
	pb "test/gen/api/v1"
)

var users []*pb.User

type Server struct {
	pb.UnimplementedBasicServiceServer
}

func NewServer() *Server {
	return &Server{
	}
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	if len(users) < 1 {
		return nil, fmt.Errorf("no users found")
	}
	return &pb.GetResponse{
		User: users,
	}, nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {

	users = append(users, &pb.User{
		Name: req.User.GetName(),
		Age: req.User.GetAge(),
	})
	
	return &pb.CreateResponse{
		Result: fmt.Sprintf("User with name: %s and age: %v was created!", req.User.GetName(), req.User.GetAge()),
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	fmt.Println("Input Name: ", req.GetName())
	for idx, user := range users{
		if user.Name == req.GetName() {
			users = append(users[:idx], users[idx+1 :]...)
			return &pb.DeleteResponse{
				User: users,
			}, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	fmt.Println("Input Name: ", req.GetName())
	for idx, user := range users{
		if user.Name == req.GetName() {
			users[idx].Name = req.GetNewName()
			return &pb.UpdateResponse{
				User: users,
			}, nil
		}
	}
	fmt.Println(users)
	return nil, fmt.Errorf("user not found")
}
