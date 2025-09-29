package repository

import (
	"context"

	"github.com/farinas09/rest-ws/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUser(ctx context.Context, id int64) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetAllUsers(ctx context.Context) (*models.Users, error)
	Close(ctx context.Context) error
}

var userImplementation UserRepository

func SetUserRepository(repository UserRepository) {
	userImplementation = repository
}

func CreateUser(ctx context.Context, user *models.User) error {
	return userImplementation.CreateUser(ctx, user)
}

func GetUser(ctx context.Context, id int64) (*models.User, error) {
	return userImplementation.GetUser(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return userImplementation.GetUserByEmail(ctx, email)
}

func GetAllUsers(ctx context.Context) (*models.Users, error) {
	return userImplementation.GetAllUsers(ctx)
}
