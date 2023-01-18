package usecase

import (
	"context"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

// TODO: value or poiner
type (
	User interface {
		SignUp(ctx context.Context, email, password *string) error
		SignIn(ctx context.Context, email, password *string) (*string, error)
		ParseToken(ctx context.Context, accessToken *string) (*entity.User, error)
	}

	UserRepo interface {
		CreateUser(ctx context.Context, user *entity.User) error
		GetUser(ctx context.Context, username, password *string) (*entity.User, error)
	}

	Position interface {
		GetAllPosition(ctx context.Context) ([]entity.Position, error)
		GetPosition(ctx context.Context) (*entity.Position, error)
		CreatePosition(ctx context.Context, position *entity.Position) error
		UpdatePosition(ctx context.Context, position *entity.Position) error
		DeletePosition(ctx context.Context, position *entity.Position) error
	}

	PositionRepo interface {
		Create(ctx context.Context, position *entity.Position) error
		FindAll(ctx context.Context) ([]entity.Position, error)
		Delete(ctx context.Context, ID *string) error
	}
)
