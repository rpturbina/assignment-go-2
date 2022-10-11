package user

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rpturbina/assigment-go-2/pkg/domain/message"
	"github.com/rpturbina/assigment-go-2/pkg/domain/user"
)

type UserHdlImpl struct {
	userUsecase user.UserUsecase
}

func (u *UserHdlImpl) GetUserByEmailHdl(ctx *gin.Context) {
	log.Printf("%T - GetUserByEmailHdl is invoked\n", u)
	defer log.Printf("%T - GetUserByEmailHdl executed\n", u)

	// get query params from url
	userEmail := ctx.Query("user_email")

	// check user email from query params, if empty -> BAD_REQUEST
	log.Println("check user email from quary params")
	if err := checkUserEmail(userEmail); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// calling service/usecase for get user data by email
	log.Println("calling get user by email service usecase")
	result, err := u.userUsecase.GetUserByEmailSvc(ctx, userEmail)
	if err != nil {
		switch err.Error() {
		case "NOT_FOUND":
			message.ErrorResponseSwitcher(ctx, http.StatusNotFound, fmt.Sprintf("user is not found: %v", userEmail))
			return
		case "INTERNAL_SERVER_ERROR":
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError)
			return
		}
	}

	ctx.JSON(http.StatusOK, message.Response{
		Code:    0,
		Message: "user is found",
		Data:    result,
	})
}

func checkUserEmail(userEmail string) error {
	if userEmail == "" {
		return errors.New("user email should not be empty")
	}
	return nil
}

func (u *UserHdlImpl) InsertUserHdl(ctx *gin.Context) {
	log.Printf("%T - InsertUserHdl is invoked\n", u)
	defer log.Printf("%T - InsertUserHdl executed\n", u)

	// binding / mendapatkan body payload dari request
	log.Println("binding body payload from request")

	var user user.User
	if err := ctx.ShouldBind(&user); err != nil {
		message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, "failed to bind payload")
		return
	}

	// checking email is empty or not, if empty => BAD_REQUEST
	log.Println("check email from request")
	if user.Email == "" {

		return
	}

	// call service/usecase for inserting the data
	log.Println("calling insert service usecase")
	result, err := u.userUsecase.InsertUserSvc(ctx, user)
	if err != nil {
		switch err.Error() {
		case "BAD_REQUEST":
			message.ErrorResponseSwitcher(ctx, http.StatusBadRequest, "invalid processing payload")
		case "INTERNAL_SERVER_ERROR":
			message.ErrorResponseSwitcher(ctx, http.StatusInternalServerError)
		}
	}

	// response result for the user if succes
	ctx.JSON(http.StatusOK, message.Response{
		Code:    0,
		Message: "success insert user",
		Data:    result,
	})
}

func NewUserHandler(userUsercase user.UserUsecase) user.UserHandler {
	return &UserHdlImpl{userUsecase: userUsercase}
}
