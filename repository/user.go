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

var implementation UserRepository

func SetRepository(repository UserRepository) {
	implementation = repository
}

func CreateUser(ctx context.Context, user *models.User) error {
	return implementation.CreateUser(ctx, user)
}

func GetUser(ctx context.Context, id int64) (*models.User, error) {
	return implementation.GetUser(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func GetAllUsers(ctx context.Context) (*models.Users, error) {
	return implementation.GetAllUsers(ctx)
}
