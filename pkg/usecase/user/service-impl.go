package user

import (
	"context"
	"errors"
	"log"

	"github.com/rpturbina/assigment-go-2/pkg/domain/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func (u *UserUsecaseImpl) GetUserByEmailSvc(ctx context.Context, email string) (result user.User, err error) {
	log.Printf("%T - GetUserByEmailSvc is invoked\n", u)
	defer log.Printf("%T - GetUserByEmailSvc executed\n", u)

	// get user from repository
	log.Println("getting user from user repository")
	result, err = u.userRepo.GetUserByEmail(ctx, email)

	if err != nil {
		// there is something wrong on the database connection
		log.Println("error when fethcing data from database: ", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")

		return result, err
	}

	// check user id > 0 ?
	log.Println("checking user id")
	if result.ID <= 0 {
		// if true, user is not found
		log.Println("user is not found: " + email)
		err = errors.New("NOT_FOUND")

		return result, err
	}
	return result, err
}

func (u *UserUsecaseImpl) InsertUserSvc(ctx context.Context, input user.User) (result user.User, err error) {
	log.Printf("%T - InsertUserSvc is invoked\n", u)
	defer log.Printf("%T - InsertUserSvc executed", u)

	// get user for input email firs
	usrCheck, err := u.GetUserByEmailSvc(ctx, input.Email)

	// check user is exist or not
	if err == nil {
		// user found
		log.Printf("user has been registered with id: %v\n", usrCheck.ID)
		err = errors.New("BAD_REQUEST")

		return result, err
	}

	// internal server error condition
	if err.Error() != "NOT_FOUND" {
		// internal server error
		log.Println("got error when checking user from database")

		return result, err
	}

	// valid condition: NOT_FOUND
	log.Println("insert user to database process")
	if err = u.userRepo.InsertUser(ctx, &input); err != nil {
		log.Printf("error when inserting user: %v\n", err.Error())
		err = errors.New("INTERNAL_SERVER_ERROR")
	}

	return input, err
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo: userRepo}
}
