package mapper

import (
	"github.com/felipematheus1337/GoBookReviewAPI/dto"
	"github.com/felipematheus1337/GoBookReviewAPI/schemas"
)

func CreateReviewToSchema(reviewDTO dto.ReviewDTO) *schemas.Review {
	return &schemas.Review{
		Reviewer: reviewDTO.Reviewer,
		Rating:   reviewDTO.Rating,
		Comment:  reviewDTO.Comment,
		BookId:   reviewDTO.BookId,
	}
}

func EntityReviewToResponse(review *schemas.Review) *schemas.ReviewResponse {
	return &schemas.ReviewResponse{
		Reviewer:  review.Reviewer,
		BookId:    review.BookId,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
		DeletedAt: review.DeletedAt,
	}
}
