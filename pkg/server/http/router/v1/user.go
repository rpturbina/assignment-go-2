package v1

import (
	"github.com/gin-gonic/gin"
	engine "github.com/rpturbina/assigment-go-2/config/gin"
	"github.com/rpturbina/assigment-go-2/pkg/domain/user"
	"github.com/rpturbina/assigment-go-2/pkg/server/http/router"
)

type UserRouterImpl struct {
	ginEngine   engine.HttpServer
	routerGroup *gin.RouterGroup
	userHandler user.UserHandler
}

func (u *UserRouterImpl) get() {
	// all path for get method are here
	u.routerGroup.GET("", u.userHandler.GetUserByEmailHdl)
}

func (u *UserRouterImpl) post() {
	// all path for post method are here
	u.routerGroup.POST("", u.userHandler.InsertUserHdl)
}

func (u *UserRouterImpl) Routers() {
	u.post()
	u.get()
}

func NewUserRouter(ginEngine engine.HttpServer, userHandler user.UserHandler) router.Router {
	routerGroup := ginEngine.GetGin().Group("/v1/user")

	return &UserRouterImpl{ginEngine: ginEngine, routerGroup: routerGroup, userHandler: userHandler}
}
