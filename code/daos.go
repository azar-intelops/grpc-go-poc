package daos

import (
	"github.com/example/example-repo/example-node/pkg/grpc/server/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var users []models.User

type UserDao struct {}

func (userDao *UserDao) CreateUser(user models.User) error {
	users = append(users, user)
	return nil
}

func (userDao *UserDao) GetUser(id string) (models.User, error) {
    if id == "" {
		return models.User{}, status.Error(codes.InvalidArgument, "id can't be empty")
	}

	for _, user := range users {
		if user.Id == id {
			return user, nil
		}
	}

	return models.User{}, status.Error(codes.Internal,"something went wrong")
}

func (userDao *UserDao) UpdateUser(id string, user models.User) error {
	if id == "" {
		return status.Error(codes.InvalidArgument, "id can't be empty")
	}

	for idx, value := range users {
		if value.Id == id {
			users[idx] = user
			return nil
		}
	}
	return status.Error(codes.Internal, "something went wrong")
}

func (userDao *UserDao) DeleteUser(id string) error {
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

func (userDao *UserDao) ListUsers() ([]models.User, error) {
	return users, nil
}
