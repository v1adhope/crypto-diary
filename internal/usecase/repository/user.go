package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgerrcode"
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

// NOTE: Return user_id???
func (ur *UserRepo) CreateUser(ctx context.Context, user *entity.User) error {
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

		return fmt.Errorf("repository: CreateUser: QueryRow: %s", err)
	}

	return nil
}

func (ur *UserRepo) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	q := `SELECT *
        FROM usr
        WHERE email = $1 AND password = $2`

	u := entity.User{}

	err := ur.Pool.QueryRow(ctx, q, username, password).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		return nil, fmt.Errorf("repository: GetUser: QueryRow: %s", err)
	}

	return &u, nil
}
