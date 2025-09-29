package repository

import (
	"context"

	"github.com/farinas09/rest-ws/models"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post *models.Post) error
	GetPostByUser(ctx context.Context, id int64) (*models.Post, error)
	GetAllPosts(ctx context.Context) (*models.Posts, error)
	DeletePost(ctx context.Context, id int64) error
	Close(ctx context.Context) error
}

var postImplementation PostRepository

func SetPostRepository(repository PostRepository) {
	postImplementation = repository
}

func CreatePost(ctx context.Context, post *models.Post) error {
	return postImplementation.CreatePost(ctx, post)
}

func GetPostByUser(ctx context.Context, id int64) (*models.Post, error) {
	return postImplementation.GetPostByUser(ctx, id)
}

func GetAllPosts(ctx context.Context) (*models.Posts, error) {
	return postImplementation.GetAllPosts(ctx)
}

func DeletePost(ctx context.Context, id int64) error {
	return postImplementation.DeletePost(ctx, id)
}

func Close(ctx context.Context) error {
	return postImplementation.Close(ctx)
}
