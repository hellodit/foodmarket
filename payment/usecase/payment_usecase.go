package usecase

import (
	"context"
	"foodmarket/domain"
	"time"
)

type paymentUsecase struct {
	PaymentRepo    domain.PaymentRepository
	OrderRepo      domain.OrderRepository
	ContextTimeout time.Duration
}

func (p paymentUsecase) Create(ctx context.Context, payment *domain.Payment) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, p.ContextTimeout)
	defer cancel()
	res, err = p.PaymentRepo.Create(ctx, payment)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewPaymentUsecase(paymentRepo domain.PaymentRepository, orderRepo domain.OrderRepository, duration time.Duration) domain.PaymentUsecase {
	return paymentUsecase{
		PaymentRepo:    paymentRepo,
		OrderRepo:      orderRepo,
		ContextTimeout: duration,
	}
}
