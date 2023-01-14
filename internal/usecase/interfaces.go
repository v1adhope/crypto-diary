package usecase

import (
	"context"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type (
	User interface {
	}

	UserRepo interface {
	}

	Position interface {
	}

	PositionRepo interface {
		Create(ctx context.Context, user *entity.User) error
		FindAll(ctx context.Context) ([]entity.User, error)
		FindOne(ctx context.Context, id string) (*entity.User, error)
		Delete(ctx context.Context) error
	}
)
