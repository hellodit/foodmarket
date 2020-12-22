package http

import (
	"context"
	"fmt"
	"foodmarket/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

type paymentHandler struct {
	paymentUsecase domain.PaymentUsecase
}

func NewPaymentHandler(e *echo.Echo, usecase domain.PaymentUsecase) {
	handler := &paymentHandler{
		paymentUsecase: usecase,
	}
	payment := e.Group("/payment")
	payment.POST("/callback", handler.PaymentCallbackHandler)
}

func (p paymentHandler) PaymentCallbackHandler(e echo.Context) error {
	ctx := e.Request().Context()
	if ctx == nil {
		ctx = context.Background()
	}

	payment := new(domain.Payment)
	if err := e.Bind(&payment); err != nil {
		return err
	}

	fmt.Printf("Payment :%s", payment)

	res, err := p.paymentUsecase.Create(ctx, payment)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}
