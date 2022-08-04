package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

//Book ...
type User struct {
	ID        string `json:"id" db:"id"`
	Email     string `json:"email"  db:"email"`
	FirstName string `json:"first_name"  db:"first_name"`
	LastName  string `json:"last_name"  db:"last_name"`
}

//UserRepository ...
type UserRepository interface {
	AddUserProfile(ctx context.Context, user User) (User, error)
	FetchUsersProfiles(ctx context.Context) (res []User, err error)
	GetUserProfileByID(ctx context.Context, id string) (User, error)
	UpdateUserProfile(ctx context.Context, user User) (User, error)
	DeleteUserProfile(ctx context.Context, id string) error
}

//UserUsecase ...
type UserUsecase interface {
	AddUserProfile(ctx context.Context, user User) (User, error)
	FetchUsersProfiles(ctx context.Context) (res []User, err error)
	GetUserProfileByID(ctx context.Context, id string) (User, error)
	UpdateUserProfile(ctx context.Context, user User) (User, error)
	DeleteUserProfile(ctx context.Context, id string) error
}

//UserUsecase ...
type UserHandler interface {
	Route(r *gin.RouterGroup)
	AddUserProfile(ctx *gin.Context)
	FetchUsersProfiles(ctx *gin.Context)
	GetUserProfileByID(ctx *gin.Context)
	UpdateUserProfile(ctx *gin.Context)
	DeleteUserProfile(ctx *gin.Context)
}
