package service

import (
	"errors"

	"github.com/felipematheus1337/GoBookReviewAPI/dto"
	_ "github.com/felipematheus1337/GoBookReviewAPI/dto"
	"github.com/felipematheus1337/GoBookReviewAPI/mapper"
	"github.com/felipematheus1337/GoBookReviewAPI/schemas"
	"gorm.io/gorm"
)

type BookService interface {
	Create(dto dto.BookDTO) (*schemas.BookResponse, error)
	List() ([]schemas.BookResponse, error)
}

type bookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) *bookService {
	return &bookService{db: db}
}

func (s *bookService) Create(bookDTO dto.BookDTO) (*schemas.BookResponse, error) {

	book := mapper.CreateToSchema(bookDTO)

	if book != nil {
		return &schemas.BookResponse{}, errors.New("Falha ao serializar DTO.")
	}

	if err := s.db.Create(&book).Error; err != nil {
		return &schemas.BookResponse{}, err
	}

	response := mapper.EntityToResponse(book)

	if response == nil {
		return &schemas.BookResponse{}, errors.New("Falha ao crear entity book response.")
	}

	return response, nil

}

func (s *bookService) List() ([]schemas.BookResponse, error) {

	var booksResponse []schemas.BookResponse

	var books []schemas.Book

	if err := s.db.Find(&books).Error; err != nil {
		return []schemas.BookResponse{}, err
	}

	for _, bookFinded := range books {
		var b *schemas.BookResponse

		b = mapper.EntityToResponse(&bookFinded)

		booksResponse = append(booksResponse, *b)
	}

	return booksResponse, nil
}
