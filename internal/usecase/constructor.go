package usecase

import "github.com/v1adhope/crypto-diary/internal/usecase/repository"

type UseCases struct {
	User User
	// Position PositionUseCase
}

func New(repos *repository.Repos) *UseCases {
	return &UseCases{
		User: NewUserCase(repos.User),
		// Position:
	}
}
