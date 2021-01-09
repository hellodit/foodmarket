package postgre

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"foodmarket/domain"
	"github.com/go-pg/pg/v10"
)

type psqlOrderRepository struct {
	DB *pg.DB
}

func (p psqlOrderRepository) FindBy(ctx context.Context, key, value string) (order *domain.Order, err error) {
	logrus.Infoln("Find Data By Key Value")
	order = new(domain.Order)
	if err := p.DB.Model(order).Where(key+"=?", value).First(); err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"data": order,
	}).Infoln("Respond Data")
	return order, nil
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
		Relation("Food").
		Limit(20).Select()

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func NewPsqlOrderRepository(db *pg.DB) domain.OrderRepository {
	return psqlOrderRepository{DB: db}
}
