package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUser(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (ur *UserRepo) CreateUser(ctx context.Context, user entity.User) error {
	q := `INSERT INTO users(email, password)
        VALUES($1,$2)`

	_, err := ur.Pool.Exec(ctx, q, user.Email, user.Password)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return entity.ErrUserAlreadyExists
			}

			return fmt.Errorf("repository: CreateUser: QyeryRow: PgErr: %s, %s", pgErr.Code, pgErr.Message)
		}

		return fmt.Errorf("repository: CreateUser: QueryRow: %w", err)
	}

	return nil
}

func (ur *UserRepo) GetUser(ctx context.Context, email string) (*entity.User, error) {
	q := `SELECT *
        FROM usr
        WHERE email = $1`

	u := &entity.User{}

	err := ur.Pool.QueryRow(ctx, q, email).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrUserNotExists
		}

		return nil, fmt.Errorf("repository: GetUser: QueryRow: %w", err)
	}

	return u, nil
}
