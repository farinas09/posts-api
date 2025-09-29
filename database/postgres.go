package database

import (
	"context"
	"database/sql"

	"github.com/farinas09/rest-ws/models"
	_ "github.com/lib/pq"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (*PostgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PostgresRepository{db}, nil
}

func (repository *PostgresRepository) CreateUser(ctx context.Context, user *models.User) error {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2) RETURNING id", user.Email, user.Password)
	err := row.Scan(&user.Id)
	return err
}

func (repository *PostgresRepository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	rows, err := repository.db.QueryContext(ctx, "SELECT id, email FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (repository *PostgresRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	rows, err := repository.db.QueryContext(ctx, "SELECT id, email, password FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	if rows.Next() {
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (repository *PostgresRepository) GetAllUsers(ctx context.Context) (*models.Users, error) {
	rows, err := repository.db.QueryContext(ctx, "SELECT id, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users models.Users
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.Id, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (repository *PostgresRepository) Close(ctx context.Context) error {
	return repository.db.Close()
}
