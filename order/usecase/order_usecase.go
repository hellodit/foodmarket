package usecase

import (
	"context"
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
	order.ID = uuid.New()
	order.Status = "pending"
	order.CreatedAt = time.Now()

	food, err := o.FoodRepo.GetByID(ctx, uuid.MustParse(form.FormValue("food_id")))
	if err != nil {
		return nil, err
	}

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
			OrderID:  order.ID.String(),
			GrossAmt: int64(order.Price),
		},
		CustomerDetail: &midtrans.CustDetail{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
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

func (o orderUsecase) SetAsPaid(ctx context.Context, OrderID uuid.UUID) (res interface{}, err error) {
	panic("implement me")
}
