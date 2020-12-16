package postgre

import (
	"context"
	"foodmarket/domain"
	"github.com/google/uuid"

	"github.com/go-pg/pg/v10"
)

type postgreFoodRepository struct {
	DB *pg.DB
}

//NewPostgreFoodRepository will create an object that represent the food.Repository interface
func NewPostgreFoodRepository(DB *pg.DB) domain.FoodRepository {
	return &postgreFoodRepository{DB}
}

func (p *postgreFoodRepository) GetByID(ctx context.Context, id uuid.UUID) (res domain.Food, err error) {
	food := new(domain.Food)
	err = p.DB.Model(food).Where("id = ?", id).Select()
	if err != nil {
		return domain.Food{}, nil
	}
	return domain.Food{}, err
}

func (p *postgreFoodRepository) Fetch(ctx context.Context) (res []domain.Food, err error) {
	var foods []domain.Food

	err = p.DB.Model(&foods).Select()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (p *postgreFoodRepository) Store(ctx context.Context, food *domain.Food) error {
	_, err := p.DB.Model(food).Insert()
	if err != nil {
		return nil
	}
	return err
}

func (p *postgreFoodRepository) GetByTitle(ctx context.Context, title string) (domain.Food, error) {
	panic("implement me")
}

func (p *postgreFoodRepository) Update(ctx context.Context, f *domain.Food) error {
	panic("implement me")
}

func (p *postgreFoodRepository) Delete(ctx context.Context, id uuid.UUID) (food *domain.Food, err error) {
	panic("implement me")
}
