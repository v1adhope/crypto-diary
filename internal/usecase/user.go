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
	hashedPassword, err := uc.hasher.GenerateHashedPassword(password)
	if err != nil {
		return fmt.Errorf("usecase: SignUp: GenerateEncryptedPassword: %w", err)
	}

	user := entity.User{
		Email:    email,
		Password: hashedPassword,
	}

	err = uc.repo.Create(ctx, user)
	if err != nil {
		return fmt.Errorf("usecase: SignUp: Create: %w", err)
	}

	return nil
}

func (uc *UserUseCase) SignIn(ctx context.Context, email, password string) (string, string, error) {
	user, err := uc.repo.Get(ctx, email)
	if err != nil {
		return "", "", fmt.Errorf("usecase: SignIn: Get %w", err)
	}

	err = uc.hasher.CompareHashAndPassword(user.Password, password)
	if err != nil {
		return "", "", fmt.Errorf("usecase: SignIn: CompareHashAndPassword %w", err)
	}

	refreshToken, accessToken, err := uc.auth.GenerateTokenPair(user.ID)
	if err != nil {
		return "", "", fmt.Errorf("usecase: SignIn: GenerateTokenPair: %w", err)
	}

	return refreshToken, accessToken, nil
}

func (uc *UserUseCase) ReissueTokens(ctx context.Context, clientToken string) (string, string, error) {
	id, err := uc.addToBlackList(ctx, clientToken)
	if err != nil {
		return "", "", fmt.Errorf("usecase: ReissueTokens: addToBlackList: %w", err)
	}

	refreshToken, accessToken, err := uc.auth.GenerateTokenPair(id)
	if err != nil {
		return "", "", fmt.Errorf("usecase: ReissueTokens: GenerateTokenPair: %w", err)
	}

	return refreshToken, accessToken, nil
}

func (uc *UserUseCase) SignOut(ctx context.Context, clientToken string) error {
	_, err := uc.addToBlackList(ctx, clientToken)
	if err != nil {
		return fmt.Errorf("usecase: SignOut: addToBlackList: %w", err)
	}

	return nil
}

func (uc *UserUseCase) addToBlackList(ctx context.Context, clientToken string) (string, error) {
	err := uc.session.CheckToken(ctx, clientToken)
	if err != nil {
		return "", fmt.Errorf("CheckToken: %w", err)
	}

	id, lifetime, err := uc.auth.ValidateRefreshToken(clientToken)
	if err != nil {
		return "", fmt.Errorf("ValidateRefreshToken: %w", err)
	}

	err = uc.session.AddToBlackList(ctx, clientToken, lifetime)
	if err != nil {
		return "", fmt.Errorf("AddToBlackList: %w", err)
	}

	return id, nil
}

func (uc *UserUseCase) CheckAuth(ctx context.Context, clientToken string) (string, error) {
	err := uc.session.CheckToken(ctx, clientToken)
	if err != nil {
		return "", fmt.Errorf("usecase: CheckAuth: CheckToken: %w", err)
	}

	id, err := uc.auth.ValidateAccessToken(clientToken)
	if err != nil {
		return "", fmt.Errorf("usecase: CheckAuth: ValidateAccessToken: %w", err)
	}

	return id, nil
}
