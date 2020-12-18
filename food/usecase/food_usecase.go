package usecase

import (
	"context"
	"foodmarket/domain"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type foodUsecase struct {
	FoodRepo       domain.FoodRepository
	ContextTimeout time.Duration
}

func (f2 *foodUsecase) GetByID(ctx context.Context, id uuid.UUID) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	food, err := f2.FoodRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (f2 *foodUsecase) Fetch(ctx context.Context, limit int) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	res, err = f2.FoodRepo.Fetch(ctx, limit)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success",
		"data":    res,
	}, nil
}

func (f2 *foodUsecase) Update(ctx context.Context, food *domain.Food, form *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	price, err := strconv.Atoi(form.FormValue("price"))
	if err != nil {
		return err
	}

	stock, err := strconv.Atoi(form.FormValue("stock"))
	if err != nil {
		return err
	}

	food.Name = form.FormValue("name")
	food.Description = form.FormValue("description")
	food.CreatedAt = time.Now()
	food.Price = price
	food.Stock = stock

	_, err = f2.FoodRepo.Update(ctx, food)

	if err != nil {
		return err
	}

	return nil

}

func (f2 *foodUsecase) Store(ctx context.Context, food *domain.Food, form *http.Request) error {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	price, err := strconv.Atoi(form.FormValue("price"))
	if err != nil {
		return err
	}

	stock, err := strconv.Atoi(form.FormValue("stock"))
	if err != nil {
		return err
	}

	food.ID = uuid.New()
	food.Name = form.FormValue("name")
	food.Description = form.FormValue("description")
	food.CreatedAt = time.Now()
	food.Price = price
	food.Stock = stock

	err = f2.FoodRepo.Store(ctx, food)
	if err != nil {
		return err
	}
	return nil
}

func (f2 *foodUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, f2.ContextTimeout)
	defer cancel()

	err := f2.FoodRepo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewFoodUsecase(repository domain.FoodRepository, duration time.Duration) domain.FoodUsecase {
	return &foodUsecase{
		FoodRepo:       repository,
		ContextTimeout: duration,
	}
}
