package controllers

import (
	"context"
	"fmt"
	pb "test/gen/api/v1"
	"test/pkg/grpc/server/services"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var userService = services.UserService{}

type Server struct {
	pb.UnimplementedBasicServiceServer
}

func NewServer() *Server {
	userService.CreateUser(&pb.User{
		Id:uuid.New().String(),
		Name:"azar",
		Age: 24,
	})
	userService.CreateUser(&pb.User{
		Id:uuid.New().String(),
		Name:"munnar",
		Age: 25,
	})
	userService.CreateUser(&pb.User{
		Id:uuid.New().String(),
		Name:"krazt",
		Age: 29,
	})
	return &Server{
	}
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	id := req.GetId()
	if id != "" {
		user,err := userService.GetUser(id)
		if err != nil{
			return nil, err
		}
		return &pb.GetResponse{
			User: user,
		}, nil
	}
	return nil, status.Error(codes.Internal, "Data not found!")
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id := uuid.New()
	user := pb.User{
		Id: id.String(),
		Name: req.GetUser().GetName(),
		Age: req.GetUser().GetAge(),
	}
	if err := userService.CreateUser(&user); err != nil {
		return nil, fmt.Errorf("unable to create the user")
	}
	
	return &pb.CreateResponse{
		Result: fmt.Sprintf("User with name: %s and age: %v was created! with id: %v", req.User.GetName(), req.User.GetAge(), user.Id),
	}, nil
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id := req.GetId()

	err := userService.DeleteUser(id)

	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{
		Message: "User Deleted Successfully!",
	}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error)  {
	err := userService.UpdateUser(req.Id, req.GetUser())
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Message: "User Updated Successfully!",
	}, nil
}

func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	users, err := userService.ListUsers()

	if err != nil {
		return nil, err
	}

	return &pb.ListResponse{
		User: users,
	}, nil
}
