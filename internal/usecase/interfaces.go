package usecase

import (
	"context"
	"time"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

// TODO: value or poiner, JWT
type (
	User interface {
		SignUp(ctx context.Context, email, password string) error
		SignIn(ctx context.Context, email, password string) (string, string, error)
		ReissueTokens(ctx context.Context, clientToken string) (string, string, error)
	}

	Position interface {
		GetAllPosition(ctx context.Context) ([]entity.Position, error)
		CreatePosition(ctx context.Context, position *entity.Position) error
		UpdatePosition(ctx context.Context, position *entity.Position) error
		DeletePosition(ctx context.Context, position *entity.Position) error
	}

	UserRepo interface {
		CreateUser(ctx context.Context, user entity.User) error
		GetUser(ctx context.Context, email string) (*entity.User, error)
	}

	PositionRepo interface {
		Create(ctx context.Context, position *entity.Position) error
		FindAll(ctx context.Context, id string) ([]entity.Position, error)
		Delete(ctx context.Context, id string) error
	}

	PasswordHasher interface {
		GenerateEncryptedPassword(password string) (string, error)
		CompareHashAndPassword(hashedPassword, password string) error
	}

	//TODO
	AuthManager interface {
		//Returning Refresh, Access tokens and error
		GenerateTokenPair(id string) (string, string, error)

		// Depending on the header "kid" selects the type of token
		ValidateToken(clientToken string) (string, time.Duration, error)
	}

	SessionStorage interface {
		AddToBlackList(ctx context.Context, token string, exp time.Duration) error
		CheckToken(ctx context.Context, token string) error
	}
)
