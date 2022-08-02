package handler

import (
	"net/http"

	domain "go-starling-middleware/domain/book"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	bookUsecase domain.BookUsecase
}

func NewBookHandler(bookUsecase domain.BookUsecase) *BookHandler {
	return &BookHandler{bookUsecase: bookUsecase}
}

func (handler *BookHandler) Route(r *gin.RouterGroup) {
	//r.Get("/hello", handler.hello)
	r.GET("/books", handler.FindBooks)
	r.GET("/books/:id", handler.FindBook)
}

func (this *BookHandler) hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}

func (this *BookHandler) FindBooks(c *gin.Context) {
	books, _ := this.bookUsecase.Fetch(c.Request.Context())
	c.JSON(200, gin.H{"data": books})
}

func (this *BookHandler) FindBook(c *gin.Context) {
	books, _ := this.bookUsecase.GetByID(c.Request.Context(), c.Param("id"))
	c.JSON(200, gin.H{"data": books})
}
