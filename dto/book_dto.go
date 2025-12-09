package dto

type BookDTO struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	PublishedYear uint16 `json:"published_year"`
}
