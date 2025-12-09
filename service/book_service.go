package service

import (
	"github.com/felipematheus1337/GoBookReviewAPI/dto"
	_ "github.com/felipematheus1337/GoBookReviewAPI/dto"
	"github.com/felipematheus1337/GoBookReviewAPI/schemas"
	"gorm.io/gorm"
)

type BookService interface {
	Create(dto dto.BookDTO) (schemas.BookResponse, error)
	List() ([]schemas.BookResponse, error)
}

type bookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *bookService {
	return &bookService{db: db}
}
