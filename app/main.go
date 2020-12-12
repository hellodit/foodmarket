package main

import (
	"foodmarket/config"
	"foodmarket/db/postgre"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	_userHttDelivery "foodmarket/user/delivery/http"
	_userPostgreRepository "foodmarket/user/repository/postgre"
	_userUseCase "foodmarket/user/usecase"
)

func main() {

	timeoutCtx := time.Duration(5) * time.Second
	config.Read()
	dbConnector := postgre.Connect()

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

	userRepo := _userPostgreRepository.NewPsqlUserRepository(dbConnector)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, timeoutCtx)
	_userHttDelivery.NewUserHandler(e, userUsecase)

	err := e.StartServer(server)
	if err != nil {
		e.Logger.Info("Shutting down the server")
	}

}
