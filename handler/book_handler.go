package handler

import (
	"github.com/felipematheus1337/GoBookReviewAPI/service"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{Service: s}
}

func (h *BookHandler) Create(ctx *gin.Context) {
	var bookDTO dto.BookDTO
}

func (h *BookHandler) List(ctx *gin.Context) {}
