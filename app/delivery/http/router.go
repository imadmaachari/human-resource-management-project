package http

import (
	controllers "go-starling-middleware/app/controller"
	"log"
	"net/http"

	container "go-starling-middleware/infrastructure"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Book *controllers.BookHandler
}

func initHandlers(di container.Container) Handlers {
	return Handlers{
		Book: controllers.NewBookHandler(di.Services.BookService),
	}
}

func StartServer(di container.Container) {
	log.Println("Initializing the Http Server at localhost:9191")
	handlers := initHandlers(di)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Rest api for ebooks managements ."})
	})

	routeGroup := r.Group("/v1")

	handlers.Book.Route(routeGroup)

	r.Run(":9191")
}
