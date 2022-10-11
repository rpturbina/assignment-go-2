package order

import (
	"context"
)

type OrderRepo interface {
	GetOrderByUser(ctx context.Context, userId uint64) (user User, err error)
	CreateOrder(ctx context.Context, inputOrder *Order) (err error)
}
