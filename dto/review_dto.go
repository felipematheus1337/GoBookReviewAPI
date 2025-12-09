package dto

type ReviewDTO struct {
	BookId   uint   `json:"book_id"`
	Reviewer string `json:"reviewer"`
	Rating   int    `json:"rating"`
	Comment  string `json:"comment"`
}
