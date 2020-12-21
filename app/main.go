package main

import (
	"foodmarket/config"
	"foodmarket/db/postgre"
	"github.com/xendit/xendit-go"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	_foodHttpDelivery "foodmarket/food/delivery/http"
	_foodPostgreRepository "foodmarket/food/repository/postgre"
	_foodUseCase "foodmarket/food/usecase"
	_orderHttpDelivery "foodmarket/order/delivery/http"
	_orderPostgreRepository "foodmarket/order/repository/postgre"
	_orderUseCase "foodmarket/order/usecase"
	_userHttDelivery "foodmarket/user/delivery/http"
	_userPostgreRepository "foodmarket/user/repository/postgre"
	_userUseCase "foodmarket/user/usecase"
)

func main() {

	timeoutCtx := time.Duration(5) * time.Second
	config.Read()
	dbConnector := postgre.Connect()
	xendit.Opt.SecretKey = viper.GetString("xendit_key")
	server := &http.Server{
		Addr:         ":" + viper.GetString("app_port"),
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to "+viper.GetString("app_name"))
	})

	orderRepo := _orderPostgreRepository.NewPsqlOrderRepository(dbConnector)
	userRepo := _userPostgreRepository.NewPsqlUserRepository(dbConnector)
	foodRepo := _foodPostgreRepository.NewPostgreFoodRepository(dbConnector)

	orderUsecase := _orderUseCase.NewOrderUsecase(orderRepo, foodRepo, timeoutCtx)
	foodUsecase := _foodUseCase.NewFoodUsecase(foodRepo, timeoutCtx)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, timeoutCtx)

	_orderHttpDelivery.NewOrderHandler(e, orderUsecase)
	_foodHttpDelivery.NewFoodHandler(e, foodUsecase)
	_userHttDelivery.NewUserHandler(e, userUsecase)

	err := e.StartServer(server)
	if err != nil {
		e.Logger.Info("Shutting down the server")
	}

}
