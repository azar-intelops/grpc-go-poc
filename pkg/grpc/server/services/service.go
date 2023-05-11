package services

import (
	pb "test/gen/api/v1"
	"test/pkg/grpc/server/daos"
)

type UserService struct{

}

var userDao = daos.UserDao{}

func (userService *UserService) CreateUser(user *pb.User) error  {
	return userDao.CreateUser(user)
}

func (userService *UserService) GetUser(id string) (*pb.User, error)  {
	return userDao.GetUser(id)
}

func (userService *UserService) ListUsers() ([]*pb.User, error)  {
	return userDao.ListUser()
}

func (userService *UserService) UpdateUser(id string, user *pb.User) (error)  {
	return userDao.UpdateUser(id, user)
}

func (userService *UserService) DeleteUser(id string) (error)  {
	return userDao.DeleteUser(id)
}