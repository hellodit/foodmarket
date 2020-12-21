package postgre

import (
	"context"
	"foodmarket/domain"
	"github.com/go-pg/pg/v10"
)

type psqlPaymentRepository struct {
	DB *pg.DB
}

func (p psqlPaymentRepository) Create(ctx context.Context, payment *domain.Payment) (res *domain.Payment, err error) {
	_, err = p.DB.Model(payment).Insert()
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (p psqlPaymentRepository) UpdateOrder(ctx context.Context, payment *domain.Payment) (res *domain.Payment, err error) {
	panic("implement me")
}

func NewPsqlPaymentRepository(db *pg.DB) domain.PaymentRepository {
	return psqlPaymentRepository{DB: db}
}
