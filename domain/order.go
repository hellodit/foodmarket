package domain

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Order struct {
	tableName struct{}  `pg:"orders"`
	ID        uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
	InvoiceID string    `pg:"invoice_id,type:varchar(255)" json:"invoice_id" form:"invoice_id"`
	Quantity  int       `pg:"quantity,type:integer(255)" json:"quantity" form:"quantity"`
	Price     int       `pg:"price,type:integer(255)" json:"price" form:"price"`
	UserID    uuid.UUID `pg:",pk" json:"user_id"`
	User      *User     `pg:"rel:has-one"`
	FoodID    uuid.UUID `pg:",pk" json:"food_id"`
	Food      *Food     `pg:"rel:has-one"`
	Status    string    `pg:"status,type:varchar(255)" json:"status" form:"status"`
	CreatedAt time.Time `pg:"default:now()" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) (res *Order, err error)
	UpdateOrder(ctx context.Context, order *Order) (res *Order, err error)
	FetchOrder(ctx context.Context, userID uuid.UUID) (res []Order, err error)
	FindBy(ctx context.Context, key, value string) (res *Order, err error)
}

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *Order, form *http.Request) (res interface{}, err error)
	FetchOrder(ctx context.Context, userID uuid.UUID) (res interface{}, err error)
	CekStatus(ctx context.Context, InvoiceID string) (res interface{}, err error)
	NotificationCallback(ctx context.Context, InvoiceID, Status string) (res interface{}, err error)
}
