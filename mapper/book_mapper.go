package mapper

import (
	"github.com/felipematheus1337/GoBookReviewAPI/dto"
	"github.com/felipematheus1337/GoBookReviewAPI/schemas"
)

func CreateToSchema(bookDTO dto.BookDTO) *schemas.Book {
	return &schemas.Book{
		Title:         bookDTO.Title,
		PublishedYear: bookDTO.PublishedYear,
		Author:        bookDTO.Author,
	}
}

func EntityToResponse(book *schemas.Book) *schemas.BookResponse {
	return &schemas.BookResponse{
		Title:         book.Title,
		PublishedYear: book.PublishedYear,
		Author:        book.Author,
		CreatedAt:     book.CreatedAt,
		UpdatedAt:     book.UpdatedAt,
		DeletedAt:     book.DeletedAt,
		Id:            book.ID,
	}
}
