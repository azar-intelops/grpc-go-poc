package services

import (
	"github.com/example/example-repo/example-node/pkg/grpc/server/daos"
	"github.com/example/example-repo/example-node/pkg/grpc/server/models"
)

type UserService struct {}

var userDao = daos.UserDao{}

func (UserService *UserService) CreateUser(user models.User) error {
	return userDao.CreateUser(user)
}

func (UserService *UserService) GetUser(id string) (models.User, error) {
	return userDao.GetUser(id)
}

func (UserService *UserService) ListUsers() ([]models.User, error) {
	return userDao.ListUsers()
}

func (UserService *UserService) UpdateUser(id string, user models.User) error {
	return userDao.UpdateUser(id, user)
}

func (UserService *UserService) DeleteUser(id string) error {
	return userDao.DeleteUser(id)
}
