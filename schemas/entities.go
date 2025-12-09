package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	title         string
	author        string
	publishedYear uint16
}

type Review struct {
	gorm.Model
	BookId   uint
	reviewer string
	rating   int
	comment  string
}

type BookResponse struct {
	Id            uint      `json:"id"`
	Title         string    `json:"title"`
	Author        string    `json:"author"`
	PublishedYear uint16    `json:"published_year"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
