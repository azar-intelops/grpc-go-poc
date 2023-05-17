package controllers

import (
	"context"
	pb "test/gen/api/v1"
	"test/pkg/grpc/server/models"
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
	userService.CreateUser(models.User{
		Id:uuid.New().String(),
		Name:"azar",
		Age: 24,
	})
	userService.CreateUser(models.User{
		Id:uuid.New().String(),
		Name:"munnar",
		Age: 25,
	})
	userService.CreateUser(models.User{
		Id:uuid.New().String(),
		Name:"krazt",
		Age: 29,
	})
	return &Server{
	}
}

func (s *Server) Ping(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{
		Message: "Server is healthy and Working!",
	}, nil
}

func (s *Server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id := uuid.New().String()
	userReq := req.GetUser()
	
	user := models.User{
		Id: id,
		Name: userReq.GetName(),
		Age: userReq.GetAge(),
	}

	if err := userService.CreateUser(user); err != nil {
		return nil, err
	}
	
	return &pb.CreateResponse{
		Result: "User Created Successfully!",
	}, nil
}

func (s *Server) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	users, err := userService.ListUsers()

	if err != nil {
		return nil, err
	}

	var userList []*pb.User

	for _, user := range users{
		userList = append(userList, &pb.User{
			Id: user.Id,
			Name: user.Name,
			Age: user.Age,
		})
	}

	return &pb.ListResponse{
		User: userList,
	}, nil
}

func (s *Server) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	id := req.GetId()
	if id != "" {
		user,err := userService.GetUser(id)
		if err != nil{
			return nil, err
		}
		return &pb.GetResponse{
			User: &pb.User{
				Id: user.Id,
				Name: user.Name,
				Age: user.Age,
			},
		}, nil
	}
	return nil, status.Error(codes.Internal, "Data not found!")
}

func (s *Server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	id := req.GetId()

	err := userService.DeleteUser(id)

	// fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteResponse{
		Message: "User Deleted Successfully!",
	}, nil
}

func (s *Server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error)  {
	request := req.GetUser()
	response := models.User{
		Id: request.Id,
		Name: request.GetName(),
		Age: request.GetAge(),
	}
	err := userService.UpdateUser(req.Id, response)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateResponse{
		Message: "User Updated Successfully!",
	}, nil
}

