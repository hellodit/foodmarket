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

	//var payment domain.Payment
	fmt.Printf("Form value %s", e.FormValue("event"))
	//res, err := p.paymentUsecase.Create(ctx, &payment, e.Request())
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusUnprocessableEntity, err.Error()).SetInternal(err)
	//}
	res := "test"
	return e.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data":   res,
	})
}
