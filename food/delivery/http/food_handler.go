package http

import (
	"context"
	_ "errors"
	"foodmarket/domain"
	"github.com/labstack/echo/v4"
	_ "github.com/thedevsaddam/govalidator"
	"net/http"
)

type foodHandler struct {
	foodUsecase domain.FoodUsecase
}

func NewFoodHandler(e *echo.Echo, usecase domain.FoodUsecase) {
	handler := &foodHandler{foodUsecase: usecase}
	e.GET("/foods", handler.FetchArticle)
	//e.POST("/food", handler.Store)
}

func (f foodHandler) FetchArticle(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, err := f.foodUsecase.Fetch(ctx)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}

//func (f foodHandler) Store(e echo.Context) error  {
//	rules := govalidator.MapData{
//		"name": []string{"required"},
//		"description": []string{"required"},
//		"stock": []string{"required"},
//		"price": []string{"required"},
//	}
//
//	validate := govalidator.Options{
//		Request:        e.Request(),
//		Rules:           rules,
//	}
//
//	err := govalidator.New(validate).Validate()
//
//	if err != nil {
//		if len(err) > 0 {
//			return echo.NewHTTPError(http.StatusUnprocessableEntity, err).SetInternal(errors.New("invalid parameter"))
//		}
//	}
//
//	ctx := e.Request().Context()
//
//
//}
