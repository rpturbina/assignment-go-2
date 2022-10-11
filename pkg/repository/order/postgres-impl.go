package order

import (
	"context"
	"log"

	"github.com/rpturbina/assigment-go-2/config/postgres"
	"github.com/rpturbina/assigment-go-2/pkg/domain/order"
)

type OrderRepoImpl struct {
	pgCln postgres.PostgresClient
}

func (o *OrderRepoImpl) GetOrderByUser(ctx context.Context, userId uint64) (result order.User, err error) {
	log.Printf("%T - GetOrderByUser is Invoked\n", o)
	defer log.Printf("%T - GetOrderByUser executed\n", o)

	db := o.pgCln.GetClient()

	db.Preload("Orders").Where("users.id = ?", userId).Find(&result)

	if err = db.Error; err != nil {
		log.Printf("error when getting order by user id %v\n", userId)
	}

	return result, err
}

func (o *OrderRepoImpl) CreateOrder(ctx context.Context, inputOrder *order.Order) (err error) {
	log.Printf("%T - CreateOrder is invoked\n", o)
	defer log.Printf("%T - CreateOrder executed\n", o)

	db := o.pgCln.GetClient()

	db.Model(&order.Order{}).Create(&inputOrder)

	if err = db.Error; err != nil {
		log.Printf("error when creating order for user id%v\n", inputOrder.UserID)
	}

	return err
}

func NewOrderRepo(pgCln postgres.PostgresClient) order.OrderRepo {
	return &OrderRepoImpl{pgCln: pgCln}
}
