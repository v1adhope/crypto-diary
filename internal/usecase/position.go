package usecase

import (
	"context"
	"fmt"

	"github.com/v1adhope/crypto-diary/internal/entity"
)

type PositionUsecase struct {
	repo PositionRepo
}

func NewPositionUseCase(repo PositionRepo) *PositionUsecase {
	return &PositionUsecase{
		repo: repo,
	}
}

func (uc *PositionUsecase) GetAll(ctx context.Context, id string) ([]entity.Position, error) {
	positions, err := uc.repo.FindAll(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("usecase: GetAll position: FindAll: %w", err)
	}

	return positions, nil
}

func (uc *PositionUsecase) Create(ctx context.Context, position *entity.Position) error {
	if err := uc.repo.Create(ctx, position); err != nil {
		return fmt.Errorf("usecase: Create position: Create: %w", err)
	}

	return nil
}

func (uc *PositionUsecase) Delete(ctx context.Context, userID, positionID string) error {
	if err := uc.repo.Delete(ctx, userID, positionID); err != nil {
		return fmt.Errorf("usecase: Delete position: Delete: %w", err)
	}

	return nil
}

func (uc *PositionUsecase) Replace(ctx context.Context, position *entity.Position) error {
	if err := uc.repo.Replace(ctx, position); err != nil {
		return fmt.Errorf("usecase: Replace position: Replace: %w", err)
	}

	return nil
}
