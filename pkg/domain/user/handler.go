package user

import "github.com/gin-gonic/gin"

type UserHandler interface {
	GetUserByEmailHdl(ctx *gin.Context)
	InsertUserHdl(ctx *gin.Context)
}
