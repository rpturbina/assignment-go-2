package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/assigment-go-2/config/gin"
	"github.com/rpturbina/assigment-go-2/pkg/domain/order"
	"github.com/rpturbina/assigment-go-2/pkg/server/http/router"
)

type OrderRouterImpl struct {
	ginEngine    engine.HttpServer
	routerGroup  *gin.RouterGroup
	orderHandler order.OrderHandler
}

func (o *OrderRouterImpl) get() {
	o.routerGroup.GET("", o.orderHandler.GetOrderByUserHdl)
}

func (o *OrderRouterImpl) post() {
	o.routerGroup.POST("", o.orderHandler.CreateOrderHdl)
}

func (o *OrderRouterImpl) Routers() {
	o.get()
	o.post()
}

func NewOrderRouter(ginEngine engine.HttpServer, orderHandler order.OrderHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/order")

	return &OrderRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, orderHandler: orderHandler}
}
