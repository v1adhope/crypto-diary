package usecase

import (
	"github.com/v1adhope/crypto-diary/internal/usecase/repository"
)

type UseCases struct {
	User User
	// Position PositionUseCase
}

type Deps struct {
	Repos   *repository.Repos
	Hasher  PasswordHasher
	Auth    AuthManager
	Session SessionStorage
}

func New(d Deps) *UseCases {
	return &UseCases{
		User: NewUserCase(d.Repos.User, d.Hasher, d.Auth, d.Session),
		// Position:
	}
}
