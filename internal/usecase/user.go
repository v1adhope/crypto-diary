package usecase

import (
	"context"
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type UserUseCase struct {
	repo    UserRepo
	hasher  PasswordHasher
	auth    AuthManager
	session SessionStorage
}

func NewUserCase(r UserRepo, h PasswordHasher, m AuthManager, s SessionStorage) *UserUseCase {
	return &UserUseCase{
		repo:    r,
		hasher:  h,
		auth:    m,
		session: s,
	}
}

func (uc *UserUseCase) SignUp(ctx context.Context, email, password string) error {
	hashedClientPassword, err := uc.hasher.GenerateEncryptedPassword(password)
	if err != nil {
		return fmt.Errorf("usecase: SignUp: GenerateEncryptedPassword: %w", err)
	}

	user := entity.User{
		Email:    email,
		Password: hashedClientPassword,
	}

	err = uc.repo.CreateUser(ctx, user)
	if err != nil {
		return fmt.Errorf("usecase: SignUp: CreateUser: %w", err)
	}

	return nil
}

func (uc *UserUseCase) SignIn(ctx context.Context, email, password string) (string, string, error) {
	user, err := uc.repo.GetUser(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("usecase: SignIn: GetUser %w", err)
	}

	err = uc.hasher.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return "", "", fmt.Errorf("usecase: SignIn: CompareHashAndPassword %w", entity.ErrWrongPassword)
	}

	refreshToken, accessToken, err := uc.auth.GenerateTokenPair(user.ID)
	if err != nil {
		return "", "", fmt.Errorf("usecase: SignIn: GenerateTokenPair: %w", err)
	}

	return refreshToken, accessToken, nil
}

func (uc *UserUseCase) ReissueTokens(ctx context.Context, clientToken string) (string, string, error) {
	err := uc.session.CheckToken(ctx, clientToken)
	if err != nil {
		return "", "", fmt.Errorf("usecase: ReissueTokens: CheckToken: %w", err)
	}

	id, lifetime, err := uc.auth.ValidateToken(clientToken)
	if err != nil {
		return "", "", fmt.Errorf("usecase: RefreshTokens: ValidateToken: %w", err)
	}

	err = uc.session.AddToBlackList(ctx, clientToken, lifetime)
	if err != nil {
		return "", "", fmt.Errorf("usecase: ReissueTokens: AddToBlackList: %w", err)
	}

	refreshToken, accessToken, err := uc.auth.GenerateTokenPair(id)
	if err != nil {
		return "", "", fmt.Errorf("usecase: ReissueTokens: GenerateTokenPair: %w", err)
	}

	return refreshToken, accessToken, nil
}
