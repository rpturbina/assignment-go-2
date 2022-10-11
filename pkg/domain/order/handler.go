package order

import "github.com/gin-gonic/gin"

type OrderHandler interface {
	GetOrderByUserHdl(ctx *gin.Context)
	CreateOrderHdl(ctx *gin.Context)
}
