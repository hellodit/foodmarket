package domain

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type (
	//Food data structure
	Food struct {
		tableName   struct{}  `pg:"foods"`
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
	Fetch(ctx context.Context, limit int) (res []Food, err error)
	Store(context.Context, *Food) error
	FindBy(ctx context.Context, key, value string) (food *Food, err error)
	Update(ctx context.Context, f *Food) (Food *Food, err error)
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (Food *Food, err error)
}

//FoodUsecase to interact with database
type FoodUsecase interface {
	Fetch(ctx context.Context, limit int) (res interface{}, err error)
	GetByID(ctx context.Context, id uuid.UUID) (res interface{}, err error)
	Update(ctx context.Context, food *Food, form *http.Request) error
	Store(ctx context.Context, food *Food, form *http.Request) error
	Delete(ctx context.Context, id uuid.UUID) error
}
