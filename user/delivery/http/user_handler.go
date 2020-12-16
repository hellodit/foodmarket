package http

import (
	"context"
	"errors"
	"foodmarket/domain"
	"foodmarket/helper"
	"foodmarket/middleware"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type userHandler struct {
	userUsecase domain.UserUseCase
}

//NewUserHandler for user
func NewUserHandler(e *echo.Echo, UserUsecase domain.UserUseCase) {
	handler := &userHandler{userUsecase: UserUsecase}
	user := e.Group("/user")
	customMiddleware := middleware.Init()
	user.GET("/:id", handler.GetByIDHandler, customMiddleware.Auth)
	user.GET("/profile", handler.ProfileHandler, customMiddleware.Auth)
	user.GET("/fetch", handler.FetchHandler, customMiddleware.Auth)
	user.POST("/register", handler.RegisterHandler)
	user.POST("/login", handler.LoginHandler)
	user.POST("/update", handler.UpdateHandler, customMiddleware.Auth)
	user.POST("/forget-password", handler.ForgerPasswordHandler, customMiddleware.Auth)
}

func (u userHandler) UpdateHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}

	var usr domain.User
	if err := e.Bind(&usr); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	claims, err := helper.ParseToken(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err).SetInternal(err)
	}

	usr.ID = uuid.MustParse(claims["sub"].(string))
	res, err := u.userUsecase.UpdateUser(ctx, &usr, e.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}
func (u userHandler) GetByIDHandler(e echo.Context) error {
	ctx := e.Request().Context()

	if ctx == nil {
		ctx = context.Background()
	}
	id, err := uuid.Parse(e.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}
	res, err := u.userUsecase.GetUserById(ctx, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

func (u userHandler) RegisterHandler(e echo.Context) error {
	rules := govalidator.MapData{
		"name":     []string{"required"},
		"password": []string{"required"},
		"email":    []string{"required"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	ctx := e.Request().Context()
	var usr domain.User

	if ctx == nil {
		ctx = context.Background()
	}

	res, err := u.userUsecase.Register(ctx, &usr, e.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

func (u userHandler) FetchHandler(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := u.userUsecase.Fetch(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusFailedDependency, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, res)
}

func (u userHandler) LoginHandler(e echo.Context) error {
	rules := govalidator.MapData{
		"password": []string{"required"},
		"email":    []string{"required"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	var credential domain.Credential

	if err := e.Bind(&credential); err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(errors.New("invalid parameter"))
	}

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := u.userUsecase.Login(ctx, &credential)

	if err != nil {
		return echo.NewHTTPError(http.StatusFailedDependency, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, res)
}

func (u userHandler) ProfileHandler(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	claims, err := helper.ParseToken(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err).SetInternal(err)
	}
	profile := claims["user"]
	return e.JSON(http.StatusOK, profile)
}

func (u userHandler) ForgerPasswordHandler(e echo.Context) error {

	rules := govalidator.MapData{
		"email": []string{"required", "email"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}

	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	res, err := u.userUsecase.ForgetPassword(ctx, e.FormValue("email"))
	if err != nil {
		return echo.NewHTTPError(res, err.Error()).SetInternal(err)

	}

	return e.JSON(http.StatusOK, map[string]interface{}{"message": "Success. Mail was sent!"})
}
