package controllers

import (
	"context"

	pb "github.com/azar-intelops/testx/assembly/gen/api/v1"
	"github.com/azar-intelops/testx/assembly/pkg/grpc/server/models"
	"github.com/azar-intelops/testx/assembly/pkg/grpc/server/services"

	"github.com/google/uuid"
)

var userService = services.UserService{}

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (s *UserServer) Ping(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: "Server is healthy and working!",
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id := uuid.New().String()
	request := req.GetUser()

	user := models.User{
		Id: id,
		Name: request.GetName(),
		Age: request.GetAge(),
		
		
	}

	if err := userService.CreateUser(user); err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Message: "User Created Successfully!",
	}, nil
}


func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	id := req.GetId()
	user, err := userService.GetUser(id)

	if err != nil {
		return nil, err
	}

	userResponse := &pb.User{
		Id: user.Id,
		Name: user.Name,
		Age: user.Age,
		
		
	}

	return &pb.GetUserResponse{
		User: userResponse,
	}, nil
}

func (s *UserServer) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, err := userService.ListUsers()

	if err != nil {
		return nil, err
	}

	var userList []*pb.User
	for _, user := range users {
		userResponse := &pb.User{
			Id: user.Id,
			Name: user.Name,
			Age: user.Age,
			
			
		}
		userList = append(userList, userResponse)
	}

	return &pb.ListUsersResponse{
		User: userList,
	}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	id := req.GetId()
	user := req.GetUser()

	userResponse := models.User{
		Id: id,
		Name: user.GetName,
		Age: user.GetAge,
		
		
	}
	err := userService.UpdateUser(id, userResponse)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserResponse{
		Message: "User Updated Successfully!",
	}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	id := req.GetId()

	err := userService.DeleteUser(id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserResponse{
		Message: "User Deleted Successfully!",
	}, nil
}
