package domain

import (
	"context"
	"gorm.io/gorm"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type (
	//Credential user
	Credential struct {
		Email    string `json:"email" form:"email"`
		Password string `json:"password" form:"password"`
	}

	//User struct
	User struct {
		ID        uuid.UUID      `json:"id"`
		Name      string         `json:"name" form:"name"`
		Email     string         `json:"email" form:"email"`
		Type      string         `json:"type" form:"type"`
		Password  string         `json:"-" form:"password"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `json:"deleted_at"`
	}
)

//UserRepository interface
type UserRepository interface {
	CreateUser(ctx context.Context, usr *User) (user *User, err error)
	Attempt(ctx context.Context, credential *Credential) (user *User, err error)
	Update(ctx context.Context, usr *User) (user *User, err error)
	Find(ctx context.Context, id uuid.UUID) (user *User, err error)
	FindBy(ctx context.Context, key, value string) (user *User, err error)
	Fetch(ctx context.Context) (res []User, err error)
}

//UserUseCase interface
type UserUseCase interface {
	Register(ctx context.Context, usr *User, form *http.Request) (res interface{}, err error)
	UpdateUser(ctx context.Context, usr *User, form *http.Request) (res interface{}, err error)
	Login(ctx context.Context, credential *Credential) (res interface{}, err error)
	Logout(ctx context.Context, claims jwt.Claims)
	GetUserById(ctx context.Context, id uuid.UUID) (res interface{}, err error)
	Fetch(ctx context.Context) (res interface{}, err error)
}
