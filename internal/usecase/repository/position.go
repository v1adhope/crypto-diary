package repository

import (
	"context"
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
	"github.com/v1adhope/crypto-diary/internal/usecase"
	"github.com/v1adhope/crypto-diary/pkg/postgres"
)

const defaultEntityCap = 25

type PositionRepo struct {
	postgres.Client
}

// TODO: VALIDATION AND CLEAN AND PASSWORD JWT
func (pr *PositionRepo) Create(ctx context.Context, user *entity.User) error {
	q := `INSERT INTO users(email, password)
        VALUES($1, $2)
        RETURNING user_id`

	err := pr.QueryRow(ctx, q, user.Email, user.Password).Scan(&user.UserID)
	if err != nil {
		return fmt.Errorf("sql request: Create: %s", err)
	}

	return nil
}

func (pr *PositionRepo) FindAll(ctx context.Context) ([]entity.User, error) {
	q := `SELECT user_id, email
        FROM users`

	rows, err := pr.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("sql request: FinAll %s", err)
	}
	defer rows.Close()

	users := make([]entity.User, 0, defaultEntityCap)

	for rows.Next() {
		u := entity.User{}

		err := rows.Scan(&u.UserID, &u.Email)
		if err != nil {
			return nil, fmt.Errorf("sql request: FinAll %s", err)
		}

		users = append(users, u)
	}

	return users, nil
}

// TODO: return * or value
func (pr *PositionRepo) FindOne(ctx context.Context, email string) (*entity.User, error) {
	q := `SELECT user_id, email
        FROM users
        WHERE email = $1`

	user := entity.User{}

	err := pr.QueryRow(ctx, q, email).Scan(&user.UserID, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("sql request: FinOne: %s", err)
	}

	return &user, nil
}

// TODO
func (pr *PositionRepo) Delete(ctx context.Context) error {
	panic("implement me")
}

func New(client postgres.Client) usecase.PositionRepo {
	return &PositionRepo{client}
}
