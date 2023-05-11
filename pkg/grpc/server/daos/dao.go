package daos

import (
	"fmt"
	pb "test/gen/api/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var users []*pb.User

type UserDao struct {

}

func (userDao *UserDao) CreateUser(user *pb.User) error  {
	users = append(users, user)
	return nil
}

func (userDao *UserDao) GetUser(id string) (*pb.User, error)  {
	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found with id: %v", id)
}

func (userDao *UserDao) UpdateUser(id string, user *pb.User) (error)  {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id can't be empty")
	}

	for idx, usr := range users {
		if usr.Id == id{
			users[idx] = user
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
	return status.Error(codes.Internal, "something went well")
}

func (userDao *UserDao) ListUser () ([]*pb.User, error) {
	if len(users) < 1 {
		return []*pb.User{}, fmt.Errorf("data is empty %v", users )
	}
	return users, nil
}