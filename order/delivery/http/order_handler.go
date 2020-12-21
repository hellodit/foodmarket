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
	"strconv"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecase
}

func NewOrderHandler(e *echo.Echo, usecase domain.OrderUsecase) {
	handler := &orderHandler{orderUsecase: usecase}
	order := e.Group("/order")
	customMiddleware := middleware.Init()
	order.POST("/create", handler.CreateOrderHandler, customMiddleware.Auth)
}

func (o orderHandler) CreateOrderHandler(e echo.Context) error {
	rules := govalidator.MapData{
		"food_id":  []string{"required"},
		"quantity": []string{"required"},
	}

	validate := govalidator.Options{
		Request: e.Request(),
		Rules:   rules,
	}

	if err := govalidator.New(validate).Validate(); len(err) > 0 {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
	}
	ctx := e.Request().Context()
	if ctx != nil {
		ctx = context.Background()
	}
	var order domain.Order

	claims, err := helper.ParseToken(e)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err).SetInternal(err)
	}

	qty, err := strconv.Atoi(e.FormValue("quantity"))

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	order.UserID = uuid.MustParse(claims["sub"].(string))
	order.FoodID = uuid.MustParse(e.FormValue("food_id"))
	order.Quantity = qty

	res, err := o.orderUsecase.CreateOrder(ctx, &order, e.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}
