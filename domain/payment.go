package domain

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Payment struct {
	tableName struct{}  `pg:"payments"`
	ID        uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
	OrderID   uuid.UUID `pg:",pk" json:"order_id"`
	Status    string    `pg:"status,type:varchar(255)" json:"status" form:"status"`
	Method    string    `pg:"method,type:varchar(255)" json:"method" form:"method"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PaymentRepository interface {
	Create(ctx context.Context, payment *Payment) (res *Payment, err error)
	UpdateOrder(ctx context.Context, payment *Payment) (res *Payment, err error)
}

type PaymentUsecase interface {
	Create(ctx context.Context, payment *Payment, form *http.Request) (res interface{}, err error)
}
