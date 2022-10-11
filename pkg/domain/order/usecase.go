package order

import (
	"context"
)

type OrderUsecase interface {
	GetOrderByUserSvc(ctx context.Context, userId uint64) (result User, err error)
	CreateOrderSvc(ctx context.Context, inputOrder Order) (result Order, err error)
}
