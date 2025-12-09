package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string
	Author        string
	PublishedYear uint16
}

type Review struct {
	gorm.Model
	BookId   uint
	Reviewer string
	Rating   int
	Comment  string
}

type BookResponse struct {
	Id            uint           `json:"id"`
	Title         string         `json:"title"`
	Author        string         `json:"author"`
	PublishedYear uint16         `json:"published_year"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deletedAt"`
}

type ReviewResponse struct {
	Id        uint           `json:"id"`
	Reviewer  string         `json:"reviewer"`
	BookId    uint           `json:"book_id"`
	Rating    int            `json:"rating"`
	Comment   string         `json:"comment"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}
