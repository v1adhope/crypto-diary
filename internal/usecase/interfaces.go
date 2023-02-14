package usecase

import (
	"context"
	"time"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type (
	User interface {
		SignUp(ctx context.Context, email, password string) error
		SignIn(ctx context.Context, email, password string) (string, string, error)
		ReissueTokens(ctx context.Context, clientToken string) (string, string, error)
		CheckAuth(ctx context.Context, clientToken string) (string, error)
	}

	Position interface {
		GetAll(ctx context.Context, id string) ([]entity.Position, error)
		Create(ctx context.Context, position *entity.Position) error
		Replace(ctx context.Context, position *entity.Position) error
		Delete(ctx context.Context, userID, positionID string) error
	}

	UserRepo interface {
		Create(ctx context.Context, user entity.User) error
		Get(ctx context.Context, email string) (*entity.User, error)
	}

	PositionRepo interface {
		Create(ctx context.Context, position *entity.Position) error
		FindAll(ctx context.Context, id string) ([]entity.Position, error)
		Delete(ctx context.Context, userID, positionID string) error
		Replace(ctx context.Context, position *entity.Position) error
	}

	PasswordHasher interface {
		GenerateHashedPassword(password string) (string, error)
		CompareHashAndPassword(hashedPassword, password string) error
	}

	AuthManager interface {
		//Returning Refresh, Access tokens and error
		GenerateTokenPair(id string) (string, string, error)

		ValidateAccessToken(clientToken string) (string, error)
		ValidateRefreshToken(clientToken string) (string, time.Duration, error)
	}

	SessionStorage interface {
		AddToBlackList(ctx context.Context, token string, exp time.Duration) error
		CheckToken(ctx context.Context, token string) error
	}
)
