package middleware

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

type Middleware struct{}

func Init() *Middleware {
	return &Middleware{}
}

func (m *Middleware) Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if !strings.Contains(tokenString, "Bearer") {
			return echo.NewHTTPError(http.StatusUnauthorized, errors.New("Token not provided"))
		}

		tokenString = strings.Replace(tokenString, "Bearer ", "", -1)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if jwt.SigningMethodHS256 != token.Method {
				return nil, errors.New("invalid token")
			}

			return []byte(viper.GetString("JWT_SECRET")), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Print(token.Claims)
			fmt.Print("===========")
			return next(c)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, err)
		}
	}
}
