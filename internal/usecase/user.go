package usecase

import (
	"context"
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type UserUseCase struct {
	repo UserRepo
}

func NewUserCase(r UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
	}
}

func (uc *UserUseCase) SignUp(ctx context.Context, email, password string) error {
	user := &entity.User{
		Email:    email,
		Password: password,
	}

	err := uc.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("usecase: SignUp: %w", err)
	}

	return nil
}

func (uc *UserUseCase) SignIn(ctx context.Context, email, password string) (*entity.User, error) {
	user, err := uc.repo.GetUser(ctx, email, password)
	if err != nil {
		return nil, fmt.Errorf("usecase: SignIn: %s", err)
	}

	return user, nil
}
