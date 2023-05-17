package daos

import (
	"fmt"
	pb "test/gen/api/v1"
	"test/pkg/grpc/server/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var users []models.User

type UserDao struct {

}

func (userDao *UserDao) CreateUser(user *pb.User) error  {
	newUser := models.User{
		Id: user.Id,
		Name: user.Name,
		Age: user.Age,
	}
	users = append(users, newUser)
	return nil
}

func (userDao *UserDao) GetUser(id string) (*pb.User, error)  {
	if id == "" {
		return nil, status.Error(codes.InvalidArgument, "id can't be empty")
	}

	for _, user := range users {
		if user.Id == id {
			return &pb.User{
				Id: user.Id,
				Name: user.Name,
				Age: user.Age,
			}, nil
		}
	}

	return nil, status.Error(codes.Internal, "something went wrong")
}

func (userDao *UserDao) UpdateUser(id string, user *pb.User) (error)  {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id can't be empty")
	}

	newUser := models.User{
		Id: user.Id,
		Name: user.Name,
		Age: user.Age,
	}

	for idx, usr := range users {
		if usr.Id == id{
			users[idx] = newUser
			return nil
		}
	}
	return status.Error(codes.Internal, "something went wrong")
}

func (userDao *UserDao) DeleteUser(id string) (error) {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id can't be empty")
	} 
	
	for idx, user := range users {
		if user.Id == id {
			users = append(users[:idx], users[idx+1:]...)
			return nil
		}
	} 
	return status.Error(codes.Internal, "something went wrong")
}

func (userDao *UserDao) ListUser () ([]*pb.User, error) {
	if len(users) < 1 {
		return []*pb.User{}, fmt.Errorf("data is empty %v", users )
	}

	var userList []*pb.User
	for _, user := range users {
		userList = append(userList, &pb.User{
			Id: user.Id,
			Name: user.Name,
			Age: user.Age,
		})
	}

	return userList, nil
}