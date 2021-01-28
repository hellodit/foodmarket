package domain

import (
	"context"
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
		tableName struct{}  `pg:"users"`
		ID        uuid.UUID `pg:"id,pk,type:uuid" json:"id"`
		Name      string    `pg:"name,type:varchar(255)" json:"name" form:"name"`
		Avatar    string    `pg:"avatar,type:varchar(255)" json:"avatar" form:"avatar"`
		Email     string    `pg:"email,type:varchar(255)" json:"email" form:"email"`
		Password  string    `pg:"password,type:varchar(255)" json:"-" form:"password"`
		Type      string    `pg:"type,type:varchar(255)" json:"type" form:"type"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Orders    []Order   `pg:"rel:has-many,array"`
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
	ForgetPassword(ctx context.Context, email string) (int, error)
}
