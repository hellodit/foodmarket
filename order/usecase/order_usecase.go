package usecase

import (
	"context"
	"fmt"
	"foodmarket/domain"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/veritrans/go-midtrans"
	"net/http"
	"time"
)

type orderUsecase struct {
	OrderRepo      domain.OrderRepository
	FoodRepo       domain.FoodRepository
	ContextTimeout time.Duration
}

func NewOrderUsecase(repository domain.OrderRepository, foodRepo domain.FoodRepository, duration time.Duration) domain.OrderUsecase {
	return orderUsecase{
		FoodRepo:       foodRepo,
		OrderRepo:      repository,
		ContextTimeout: duration,
	}
}

func (o orderUsecase) CreateOrder(ctx context.Context, order *domain.Order, form *http.Request) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, o.ContextTimeout)
	defer cancel()
	food, err := o.FoodRepo.GetByID(ctx, uuid.MustParse(form.FormValue("food_id")))
	if err != nil {
		return nil, err
	}

	order.ID = uuid.New()
	order.Status = "pending"
	order.CreatedAt = time.Now()
	order.InvoiceID = Uniqid("INV")
	order.Price = food.Price * order.Quantity
	res, err = o.OrderRepo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	midclient := midtrans.NewClient()
	midclient.ServerKey = viper.GetString("midtrans_server_key")
	midclient.ClientKey = viper.GetString("midtrans_client_key")
	midclient.APIEnvType = midtrans.Sandbox
	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  order.InvoiceID,
			GrossAmt: int64(order.Price),
		},
		Items: &[]midtrans.ItemDetail{
			{
				ID:    food.ID.String(),
				Price: int64(food.Price),
				Qty:   int32(order.Quantity),
				Name:  food.Name,
			},
		},
	}

	snapTokenResp, err := snapGateway.GetToken(snapReq)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"order":    res,
		"midtrans": snapTokenResp,
	}, nil
}

func Uniqid(prefix string) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	return fmt.Sprintf("%s-%08x%05x", prefix, sec, usec)
}

func (o orderUsecase) FetchOrder(ctx context.Context, userID uuid.UUID) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, o.ContextTimeout)
	defer cancel()

	res, err = o.OrderRepo.FetchOrder(ctx, userID)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"message": "success",
		"data":    res,
	}, nil
}

func (o orderUsecase) CekStatus(ctx context.Context, InvoiceID string) (res interface{}, err error) {
	panic("implement me")
}

func (o orderUsecase) NotificationCallback(ctx context.Context, InvoiceID, orderStatusStr string) (res interface{}, err error) {
	ctx, cancel := context.WithTimeout(ctx, o.ContextTimeout)
	defer cancel()
	order, err := o.OrderRepo.FindBy(ctx, "invoice_id", InvoiceID)
	if err != nil {
		return nil, err
	}

	order.Status = orderStatusStr

	res, err = o.OrderRepo.UpdateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return res, nil
}
