package usecase

import (
	"context"
	"foodmarket/domain"
	"net/http"
	"time"
)

type paymentUsecase struct {
	PaymentRepo    domain.PaymentRepository
	OrderRepo      domain.OrderRepository
	ContextTimeout time.Duration
}

func (p paymentUsecase) Create(ctx context.Context, payment *domain.Payment, form *http.Request) (res interface{}, err error) {

	panic("implement me")
}

func NewPaymentUsecase(paymentRepo domain.PaymentRepository, orderRepo domain.OrderRepository, duration time.Duration) domain.PaymentUsecase {
	return paymentUsecase{
		PaymentRepo:    paymentRepo,
		OrderRepo:      orderRepo,
		ContextTimeout: duration,
	}
}
