package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Payment struct {
	tableName   struct{}  `pg:"payments"`
	ID          uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
	Event       string    `pg:"event,type:varchar(255)" json:"event" form:"event"`
	Phone       string    `pg:"phone,type:varchar(255)" json:"phone" form:"phone"`
	Amount      float64   `pg:"amount,type:varchar(255)" json:"amount" form:"amount"`
	Status      string    `pg:"status,type:varchar(255)" json:"status" form:"status"`
	BusinessID  string    `pg:"business_id,type:varchar(255)" json:"business_id" form:"business_id"`
	ExternalID  string    `pg:"external_id,type:varchar(255)" json:"external_id" form:"external_id"`
	EwalletType string    `pg:"ewallet_type,type:varchar(255)" json:"ewallet_type" form:"ewallet_type"`
	CreatedAt   time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) (res *Payment, err error)
	UpdateOrder(ctx context.Context, payment *Payment) (res *Payment, err error)
}

type PaymentUsecase interface {
	Create(ctx context.Context, payment *Payment) (res interface{}, err error)
}
