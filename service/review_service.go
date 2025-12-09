package service

import "gorm.io/gorm"

type ReviewService interface {
}

type reviewService struct {
	db *gorm.DB
}

func NewReviewService(db *gorm.DB) ReviewService {
	return &reviewService{db: db}
}
