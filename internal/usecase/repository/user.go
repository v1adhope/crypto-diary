package repository

import (
	"context"
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func (ur *UserRepo) CreateUser(ctx context.Context, user *entity.User) error {
	q := `INSERT INTO users(email, password)
        VALUES($1,$2)
        RETURNING user_id`

	err := ur.Pool.QueryRow(ctx, q, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("sql request: CreateUser: QueryRow: %s", err)
	}

	return nil
}

func (ur *UserRepo) GetUser(ctx context.Context, username, password *string) (*entity.User, error) {
	q := `SELECT *
        FROM usr
        WHERE email = $1 AND password = $2`

	u := entity.User{}
	err := ur.Pool.QueryRow(ctx, q, username, password).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, fmt.Errorf("sql request: GetUser: QueryRow: %s", err)
	}

	return &u, nil
}

func NewUser(pg *postgres.Postgres) usecase.UserRepo {
	return &UserRepo{pg}
}
