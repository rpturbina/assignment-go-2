package message

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code      int        `json:"code"`
	Message   string     `json:"message"`
	Error     string     `json:"error,omitempty"`      // nullable
	Data      any        `json:"data,omitempty"`       // nullable
	StartTime *time.Time `json:"start_time,omitempty"` // nullable
}

func ErrorResponseSwitcher(ctx *gin.Context, httpCode int, errorMessage ...string) {
	var response Response
	switch httpCode {
	case http.StatusNotFound:
		response = Response{
			Code:  80,
			Error: errorMessage[0],
		}
	case http.StatusBadRequest:
		response = Response{
			Code:  80,
			Error: errorMessage[0],
		}
	case http.StatusUnauthorized:
		response = Response{
			Code:    98,
			Message: "unauthorized request",
		}
	case http.StatusInternalServerError:
		response = Response{
			Code:    99,
			Message: "something went wrong",
			Error:   "INTERNAL_SERVER_ERROR",
		}
	}
	ctx.AbortWithStatusJSON(httpCode, response)
}
