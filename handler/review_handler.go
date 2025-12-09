package handler

import (
	"net/http"

	"github.com/felipematheus1337/GoBookReviewAPI/dto"
	"github.com/felipematheus1337/GoBookReviewAPI/service"
	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	Service service.ReviewService
}

func NewReviewHandler(s service.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: s}
}

func (h *ReviewHandler) Create(ctx *gin.Context) {
	var reviewDTO dto.ReviewDTO

	if err := ctx.ShouldBind(&reviewDTO); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.Service.Create(reviewDTO)

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "create-review", response, http.StatusOK)
}

func (h *ReviewHandler) List(ctx *gin.Context) {
	reviews, err := h.Service.List()

	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "list-reviews", reviews, http.StatusOK)

}
