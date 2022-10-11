package user

import "context"

type UserRepo interface {
	GetUserByEmail(ctx context.Context, email string) (user User, err error)
	InsertUser(ctx context.Context, user *User) (err error)
}
