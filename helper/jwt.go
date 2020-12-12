package helper

import (
	"context"
	"foodmarket/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"strings"
	"time"
)

type Claims struct {
	jwt.StandardClaims
	User *domain.User `json:"user"`
}

func GenerateJwt(ctx context.Context, user *domain.User) (token string, exp int64, err error){
	secret := viper.GetString("JWT_SECRET")
	tkExp := viper.GetDuration("JWT_EXPIRED_TOKEN_DURATION") * time.Minute
	jwtid := uuid.New()

	claims := Claims{
		StandardClaims : jwt.StandardClaims{
			Id:        jwtid.String(),
			Issuer:    "foodmarket",
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Add(tkExp).Unix(),
			Subject:   user.ID.String(),
		},
		User: user,
	}

	tokenClaims := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	token, err = tokenClaims.SignedString([]byte(secret))
	if err != nil {
		return "", 0, err
	}

	return token, claims.ExpiresAt, nil

}

func ValidateJwt(c echo.Context)(jwt.MapClaims, error) {
	tokenString := c.Request().Header.Get("Authorization")
	if strings.Contains(tokenString, "Bearer"){
		return nil, errors.New("Token not provided")
	}
	tokenString = strings.Replace(tokenString, "Bearer ", "", -1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.SigningMethodHS256 != token.Method {
			return nil, errors.New("invalid token")
		}

		return []byte(viper.GetString("JWT_SECRET")), nil
	})

	if err != nil{
		return nil,err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return claims, nil
	}

	return claims, err
}
