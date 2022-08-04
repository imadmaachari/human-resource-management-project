package handler

import (
	domain "user-management-project/app/domain"

	"github.com/gin-gonic/gin"
)

type UserHandlerImpl struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) domain.UserHandler {
	return &UserHandlerImpl{userUsecase: userUsecase}
}

func (handler *UserHandlerImpl) Route(r *gin.RouterGroup) {
	r.GET("/users", handler.FetchUsersProfiles)
	r.POST("/users", handler.AddUserProfile)
	r.GET("/users/:id", handler.GetUserProfileByID)
	r.PUT("/users", handler.UpdateUserProfile)
	r.DELETE("/users/:id", handler.DeleteUserProfile)
}

func (this *UserHandlerImpl) AddUserProfile(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	addedUser, err := this.userUsecase.AddUserProfile(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"data": addedUser})
}
func (this *UserHandlerImpl) FetchUsersProfiles(ctx *gin.Context) {
	users, err := this.userUsecase.FetchUsersProfiles(ctx.Request.Context())
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"data": users})
}
func (this *UserHandlerImpl) GetUserProfileByID(ctx *gin.Context) {
	user, err := this.userUsecase.GetUserProfileByID(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"data": user})
}
func (this *UserHandlerImpl) UpdateUserProfile(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	_, err := this.userUsecase.UpdateUserProfile(ctx.Request.Context(), user)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"data": "User updated successfully ."})
}
func (this *UserHandlerImpl) DeleteUserProfile(ctx *gin.Context) {
	err := this.userUsecase.DeleteUserProfile(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully."})
}
