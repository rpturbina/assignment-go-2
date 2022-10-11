package order

import (
	"context"
	"errors"
	"log"

	"github.com/rpturbina/assigment-go-2/pkg/domain/order"
)

type OrderUsecaseImpl struct {
	orderRepo order.OrderRepo
}

func (o *OrderUsecaseImpl) GetOrderByUserSvc(ctx context.Context, userId uint64) (result order.User, err error) {
	log.Printf("%T - GetOrderByUserSvc is invoked\n", o)
	defer log.Printf("%T - GetOrderByUserSvc executed\n", o)

	log.Println("getting order from order repository")
	result, err = o.orderRepo.GetOrderByUser(ctx, userId)

	if err != nil {
		log.Println("error when fetching data from database: ", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")

		return result, err
	}

	log.Println("checking order")
	if result.ID <= 0 {
		// log.Printf("user is not found: %v", email)
		err = errors.New("NOT_FOUND")

		return result, err
	}

	return result, err
}

func (o *OrderUsecaseImpl) CreateOrderSvc(ctx context.Context, inputOrder order.Order) (result order.Order, err error) {
	log.Printf("%T - CreateOrderSvc is invoked\n", o)
	defer log.Printf("%T - CreateOrderSvc executed", o)

	log.Println("create order to database process")
	if err = o.orderRepo.CreateOrder(ctx, &inputOrder); err != nil {
		log.Printf("error when inserting user: %v\n", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}

	return inputOrder, err
}

func NewOrderUsecase(orderRepo order.OrderRepo) order.OrderUsecase {
	return &OrderUsecaseImpl{orderRepo: orderRepo}
}
