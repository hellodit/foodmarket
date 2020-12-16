package usecase

import (
	"context"
	"foodmarket/domain"
	"foodmarket/helper"
	"foodmarket/thirdparty/mail"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/gommon/random"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

type UserUsecase struct {
	UserRepo       domain.UserRepository
	ContextTimeout time.Duration
}

func (u UserUsecase) ForgetPassword(ctx context.Context, email string) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	userAccount, err := u.UserRepo.FindBy(ctx, "email", email)
	if err != nil {
		return 420, err
	}

	randomPass := random.String(8)
	encrypted, _ := bcrypt.GenerateFromPassword([]byte(randomPass), bcrypt.DefaultCost)
	userAccount.Password = string(encrypted)
	_, err = u.UserRepo.Update(ctx, userAccount)
	if err != nil {
		return 423, err
	}

	subject := "Mail Subject"
	message := "Password updated to " + randomPass

	err = mail.SendSmtpMail(subject, email, message)

	if err != nil {
		return 500, err
	}

	return 200, nil

}

func (u UserUsecase) Fetch(ctx context.Context) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	res, err = u.UserRepo.Fetch(ctx)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success",
		"data":    res,
	}, nil

}

func (u UserUsecase) Register(ctx context.Context, usr *domain.User, form *http.Request) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.FormValue("password")), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	usr.ID = uuid.New()
	usr.Name = form.FormValue("name")
	usr.Email = form.FormValue("email")
	usr.CreatedAt = time.Now()
	usr.Type = "user"
	usr.Password = string(hashedPassword)

	user, err := u.UserRepo.CreateUser(ctx, usr)

	if err != nil {
		return
	}

	return user, nil
}

func (u UserUsecase) UpdateUser(ctx context.Context, usr *domain.User, form *http.Request) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	usr.Name = form.FormValue("name")
	usr.Email = form.FormValue("email")
	usr.Type = "user"
	usr.UpdatedAt = time.Now()
	usr.Password = ""

	res, err = u.UserRepo.Update(ctx, usr)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u UserUsecase) Login(ctx context.Context, credential *domain.Credential) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	user, err := u.UserRepo.Attempt(ctx, credential)
	if err != nil {
		return nil, err
	}

	token, exp, err := helper.GenerateJwt(ctx, user)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token_type":   "Bearer",
		"access_token": token,
		"expires_in":   exp,
		"profile":      user,
	}, nil
}

func (u UserUsecase) Logout(ctx context.Context, claims jwt.Claims) {
	panic("implement me")
}

func (u UserUsecase) GetUserById(ctx context.Context, id uuid.UUID) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, u.ContextTimeout)
	defer cancel()

	user, err := u.UserRepo.Find(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func NewUserUsecase(repository domain.UserRepository, duration time.Duration) domain.UserUseCase {
	return UserUsecase{
		UserRepo:       repository,
		ContextTimeout: duration,
	}

}
