package service

import "gorm.io/gorm"

type BookService interface {
}

type bookService struct {
	db *gorm.DB
}

func NewBookService(db *gorm.DB) BookService {
	return &bookService{db: db}
}
