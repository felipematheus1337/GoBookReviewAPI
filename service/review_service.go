package service

import (
	"errors"

	"github.com/felipematheus1337/GoBookReviewAPI/dto"
	"github.com/felipematheus1337/GoBookReviewAPI/mapper"
	"github.com/felipematheus1337/GoBookReviewAPI/schemas"
	"gorm.io/gorm"
)

type ReviewService interface {
	Create(reviewDTO dto.ReviewDTO) (schemas.ReviewResponse, error)
	List() ([]schemas.ReviewResponse, error)
}

type reviewService struct {
	db *gorm.DB
}

func NewReviewService(db *gorm.DB) *reviewService {
	return &reviewService{db: db}
}

func (s *reviewService) Create(reviewDTO dto.ReviewDTO) (*schemas.ReviewResponse, error) {
	schemaReview := mapper.CreateReviewToSchema(reviewDTO)

	if err := s.db.Create(&schemaReview).Error; err != nil {
		return &schemas.ReviewResponse{}, errors.New("Error creating review")
	}

	response := mapper.EntityReviewToResponse(schemaReview)

	if response == nil {
		return &schemas.ReviewResponse{}, errors.New("Error serializing to review response.")
	}

	return response, nil

}

func (s *reviewService) List() (*[]schemas.ReviewResponse, error) {

	var reviews []schemas.Review

	var reviewsResponse []schemas.ReviewResponse

	if err := s.db.Find(&reviews).Error; err != nil {
		return &[]schemas.ReviewResponse{}, errors.New("Error getting reviews")
	}

	for _, review := range reviews {
		var r *schemas.ReviewResponse

		r = mapper.EntityReviewToResponse(&review)

		reviewsResponse = append(reviewsResponse, *r)

	}

	return &reviewsResponse, nil

}
