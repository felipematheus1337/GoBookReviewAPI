package handler

import (
	"net/http"

	"github.com/felipematheus1337/GoBookReviewAPI/dto"
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

	if err := ctx.ShouldBind(&bookDTO); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.Service.Create(bookDTO)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-book", response, http.StatusCreated)
}

func (h *BookHandler) List(ctx *gin.Context) {}
