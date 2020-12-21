package usecase

import (
	"context"
	"foodmarket/domain"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type orderUsecase struct {
	OrderRepo      domain.OrderRepository
	FoodRepo       domain.FoodRepository
	ContextTimeout time.Duration
}

func (o orderUsecase) CreateOrder(ctx context.Context, order *domain.Order, form *http.Request) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, o.ContextTimeout)
	defer cancel()
	order.ID = uuid.New()
	order.Status = "pending"
	order.CreatedAt = time.Now()

	food, err := o.FoodRepo.GetByID(ctx, uuid.MustParse(form.FormValue("food_id")))
	if err != nil {
		return nil, err
	}

	order.Price = food.Price * order.Quantity

	res, err = o.OrderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (o orderUsecase) FetchOrder(ctx context.Context, userID uuid.UUID) (res interface{}, err error) {
	panic("implement me")
}

func (o orderUsecase) SetAsPaid(ctx context.Context, OrderID uuid.UUID) (res interface{}, err error) {
	panic("implement me")
}

func NewOrderUsecase(repository domain.OrderRepository, duration time.Duration) domain.OrderUsecase {
	return orderUsecase{
		OrderRepo:      repository,
		ContextTimeout: duration,
	}
}
