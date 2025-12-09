package schemas

import "gorm.io/gorm"

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
