package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type (
	//Food data structure
	Food struct {
		ID          uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
		Name        string    `pg:"name,type:varchar(255)" json:"name" form:"name"`
		Description string    `pg:"description,type:text" json:"description" form:"description"`
		Stock       int       `pg:"stock,type:integer(255)" json:"stock" form:"stock"`
		Price       int       `pg:"price,type:integer(255)" json:"price" form:"price"`
		CreatedAt   time.Time `pg:"default:now()" json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)

//FoodRepository to interact with database
type FoodRepository interface {
	Fetch(ctx context.Context) (res []Food, err error)
	Store(context.Context, *Food) error
	GetByTitle(ctx context.Context, title string) (Food, error)
	Update(ctx context.Context, f *Food) error
	GetByID(ctx context.Context, id uuid.UUID) (Food, error)
	Delete(ctx context.Context, id uuid.UUID) (food *Food, err error)
}

//FoodUsecase to interact with database
type FoodUsecase interface {
	Fetch(ctx context.Context) (res []Food, err error)
	GetByID(ctx context.Context, id uuid.UUID) (Food, error)
	GetByTitle(ctx context.Context, title string) (Food, error)
	Update(ctx context.Context, food *Food) error
	Store(ctx context.Context, food *Food) error
	Delete(ctx context.Context, id uuid.UUID) error
}
