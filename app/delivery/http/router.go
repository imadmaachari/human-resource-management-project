package http

import (
	"log"
	"net/http"
	controllers "user-management-project/app/controller"
	domain "user-management-project/app/domain"

	container "user-management-project/infrastructure"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User domain.UserHandler
}

func initHandlers(di container.Container) Handlers {
	return Handlers{
		User: controllers.NewUserHandler(di.Services.UserService),
	}
}

func StartServer(di container.Container) {
	log.Println("Initializing the Http Server at localhost:9191")
	handlers := initHandlers(di)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Rest api for HR managements ."})
	})

	routeGroup := r.Group("/v1")

	handlers.User.Route(routeGroup)

	r.Run(":8080")
}
