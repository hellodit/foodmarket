package postgre

import (
	"context"
	"github.com/google/uuid"

	"foodmarket/domain"
	"github.com/go-pg/pg/v10"
)

type psqlOrderRepository struct {
	DB *pg.DB
}

func (p psqlOrderRepository) CreateOrder(ctx context.Context, order *domain.Order) (res *domain.Order, err error) {
	_, err = p.DB.Model(order).Insert()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (p psqlOrderRepository) UpdateOrder(ctx context.Context, order *domain.Order) (res *domain.Order, err error) {
	_, err = p.DB.Model(order).
		WherePK().
		Update()

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (p psqlOrderRepository) FetchOrder(ctx context.Context, userID uuid.UUID) (res []domain.Order, err error) {
	var orders []domain.Order
	err = p.DB.Model(&orders).Where("user_id = ? ", userID).
		Order("created_at ASC").
		Limit(20).Select()

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func NewPsqlOrderRepository(db *pg.DB) domain.OrderRepository {
	return psqlOrderRepository{DB: db}
}
