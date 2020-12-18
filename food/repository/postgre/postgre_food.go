package postgre

import (
	"context"
	"foodmarket/domain"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type postgreFoodRepository struct {
	DB *pg.DB
}

//NewPostgreFoodRepository will create an object that represent the food.Repository interface
func NewPostgreFoodRepository(DB *pg.DB) domain.FoodRepository {
	return &postgreFoodRepository{DB}
}

func (p *postgreFoodRepository) GetByID(ctx context.Context, id uuid.UUID) (res *domain.Food, err error) {
	food := new(domain.Food)
	err = p.DB.Model(food).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return food, nil
}

func (p *postgreFoodRepository) Fetch(ctx context.Context, limit int) (res []domain.Food, err error) {
	var foods []domain.Food
	if limit == 0 {
		limit = 10
	}

	err = p.DB.Model(&foods).
		Order("created_at ASC").
		Limit(limit).Select()

	if err != nil {
		return nil, err
	}

	return foods, nil
}

func (p *postgreFoodRepository) Store(ctx context.Context, food *domain.Food) error {
	_, err := p.DB.Model(food).Insert()
	if err != nil {
		return nil
	}
	return err
}
func (p *postgreFoodRepository) FindBy(ctx context.Context, key, value string) (food *domain.Food, err error) {
	logrus.Infoln("Find Data By Key Value")
	food = new(domain.Food)
	if err := p.DB.Model(food).Where(key+"=?", value).First(); err != nil {
		return nil, err
	}
	logrus.WithFields(logrus.Fields{
		"data": food,
	}).Infoln("Respond Data")
	return food, nil
}

func (p *postgreFoodRepository) Update(ctx context.Context, f *domain.Food) (food *domain.Food, err error) {
	_, err = p.DB.Model(f).WherePK().Update()
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (p *postgreFoodRepository) Delete(ctx context.Context, id uuid.UUID) error {
	food := new(domain.Food)
	_, err := p.DB.Model(food).Where("id = ?", id).Delete()
	if err != nil {
		return err
	}
	return nil

}
