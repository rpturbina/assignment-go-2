package order

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/assigment-go-2/pkg/domain/message"
	"github.com/rpturbina/assigment-go-2/pkg/domain/order"
)

type OrderHdlImpl struct {
	orderUsecase order.OrderUsecase
}

func (o *OrderHdlImpl) GetOrderByUserHdl(ctx *gin.Context) {
	log.Printf("%T - GetOrderByUserHdl is invoked\n", o)
	defer log.Printf("%T - GetOrderByUserHdl executed\n", o)

	userId := ctx.Query("user_id")

	// check user id from query params, if empty -> BAD_REQUEST
	log.Println("check user id from quary params")
	if err := checkUserId(userId); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// convert string to int
	userIdInt, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError)
		log.Println(err)
		return
	}

	// calling service/usecase for get order data by user id
	log.Println("calling get order by user service usecase")
	result, err := o.orderUsecase.GetOrderByUserSvc(ctx, userIdInt)
	if err != nil {
		switch err.Error() {
		case "NOT_FOUND":
			message.ErrorResponseSwitcher(ctx, http.StatusNotFound, fmt.Sprintf("user is not found: %v", userId))
			return
		case "INTERNAL_SERVER_ERROR":
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, message.Response{
		Code:    0,
		Message: fmt.Sprintf("order by user id %v is found", userId),
		Data:    result,
	})
}

func checkUserId(userId string) error {
	if userId == "" {
		return errors.New("user id should not be empty")
	}
	return nil
}

func (o *OrderHdlImpl) CreateOrderHdl(ctx *gin.Context) {
	log.Printf("%T - CreateOrderHdl is invoked\n", o)
	defer log.Printf("%T - CreateOrderHdl executed\n", o)

	log.Println("binding body payload from request")

	var order order.Order

	if err := ctx.ShouldBind(&order); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, "failed to bind payload")
		log.Println(err)
		return
	}

	// checking item is empty or not, if empty => BAD_REQUEST
	log.Println("check item from request")
	if order.Item == "" {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, "item order should not be empty")
		return
	}

	// call service/usecase for creating the order
	log.Println("calling create order service usecase")
	result, err := o.orderUsecase.CreateOrderSvc(ctx, order)

	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, "invalid processing payload")
			return
		case "INTERNAL_SERVER_ERROR":
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusCreated, message.Response{
		Code:    0,
		Message: "success create order",
		Data:    result,
	})
}

func NewOrderHandler(orderUsecase order.OrderUsecase) order.OrderHandler {
	return &OrderHdlImpl{orderUsecase: orderUsecase}
}
