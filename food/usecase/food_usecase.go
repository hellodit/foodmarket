package usecase

import (
	"context"
	"foodmarket/domain"
	"github.com/google/uuid"
	"time"
)

type foodUsecase struct {
	FoodRepo       domain.FoodRepository
	ContextTimeout time.Duration
}

func (f2 *foodUsecase) Fetch(ctx context.Context) (res []domain.Food, err error) {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	res, err = f2.FoodRepo.Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return
}

func (f2 *foodUsecase) GetByID(ctx context.Context, id uuid.UUID) (domain.Food, error) {
	panic("implement me")
}

func (f2 *foodUsecase) GetByTitle(ctx context.Context, title string) (domain.Food, error) {
	panic("implement me")
}

func (f2 *foodUsecase) Update(ctx context.Context, food *domain.Food) error {
	panic("implement me")
}

func (f2 *foodUsecase) Store(ctx context.Context, food *domain.Food) error {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	err := f2.FoodRepo.Store(ctx, food)
	return err
}

func (f2 *foodUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}

func NewFoodUsecase(repository domain.FoodRepository, duration time.Duration) domain.FoodUsecase {
	return &foodUsecase{
		FoodRepo:       repository,
		ContextTimeout: duration,
	}
}
