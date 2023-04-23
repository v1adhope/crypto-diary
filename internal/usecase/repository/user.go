// TODO: Outside Dependencies
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

func (ur *UserRepo) Create(ctx context.Context, user entity.User) error {
	sql, args, err := ur.Builder.Insert("users").
		Columns("email", "password").
		Values(user.Email, user.Password).
		ToSql()
	if err != nil {
		return fmt.Errorf("repository: Create user: Query builder: %w", err)
	}

	_, err = ur.Pool.Exec(ctx, sql, args...)
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) {
			if pgErr.Code == pgerrcode.UniqueViolation {
				return entity.ErrUserAlreadyExists
			}
		}

		return fmt.Errorf("repository: Create user: QueryRow: %w", err)
	}

	return nil
}

func (ur *UserRepo) Get(ctx context.Context, email string) (*entity.User, error) {
	sql, args, err := ur.Builder.Select("*").
		From("get_user").
		Where("email = ?", email).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("repository: Create user: Query builder: %w", err)
	}

	u := &entity.User{}

	err = ur.Pool.QueryRow(ctx, sql, args...).Scan(&u.ID, &u.Email, &u.Password)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, entity.ErrUserNotExists
		}

		return nil, fmt.Errorf("repository: Get user: QueryRow: %w", err)
	}

	return u, nil
}
