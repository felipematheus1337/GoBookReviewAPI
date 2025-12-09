package handler

import "github.com/gin-gonic/gin"

type ReviewHandler struct {
	Service service.ReviewService
}

func NewReviewHandler(s service.ReviewService) *ReviewHandler {
	return &ReviewHandler{Service: s}
}

func (h *ReviewHandler) Create(ctx *gin.Context) {}

func (h *ReviewHandler) List(ctx *gin.Context) {}
