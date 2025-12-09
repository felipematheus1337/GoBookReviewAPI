package handler

import (
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Service service.BookService
}

func NewBookHandler(s service.BookService) *BookHandler {
	return &BookHandler{Service: s}
}

func (h *BookHandler) Create(ctx *gin.Context) {}

func (h *BookHandler) List(ctx *gin.Context) {}
