package http

import (
	"context"
	"errors"
	"foodmarket/domain"
	"github.com/labstack/echo/v4"
	"github.com/thedevsaddam/govalidator"
	"net/http"
	"strconv"
)

type foodHandler struct {
	foodUsecase domain.FoodUsecase
}

func NewFoodHandler(e *echo.Echo, usecase domain.FoodUsecase) {
	handler := &foodHandler{foodUsecase: usecase}
	e.GET("/foods", handler.FetchFoods)
	e.POST("/food", handler.StoreHandler)
	//e.POST("/food", handler.Store)
}
func (f foodHandler) FetchFoods(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}
	i, err := strconv.Atoi(e.QueryParam("limit"))

	if err != nil {
		i = 10
	}

	res, err := f.foodUsecase.Fetch(ctx, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, res)

}

func (f foodHandler) StoreHandler(e echo.Context) error {
	var food domain.Food
	rules := govalidator.MapData{
		"name":        []string{"required"},
		"description": []string{"required"},
		"stock":       []string{"required"},
		"price":       []string{"required"},
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

	err := f.foodUsecase.Store(ctx, &food, e.Request())

	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   food,
	})

}
