package main

import (
	"fmt"
	"os"

	"github.com/felipematheus1337/GoBookReviewAPI/config"
	"github.com/felipematheus1337/GoBookReviewAPI/handler"
	"github.com/felipematheus1337/GoBookReviewAPI/router"
	"github.com/felipematheus1337/GoBookReviewAPI/service"

	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Init()

	if err != nil {
		fmt.Printf("Error initializing config: %v", err)
	}

	r := gin.Default()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080"
	}

	db, err := config.InitializeMySQL()

	if err != nil {
		fmt.Printf("Error initializing MySQL: %v", err)
	}

	bookService := service.NewBookService(db)
	reviewService := service.NewReviewService(db)

	bookHandler := handler.NewBookHandler(bookService)
	reviewHandler := handler.NewReviewHandler(reviewService)

	router.InitializeRoutes(r, bookHandler, reviewHandler)

	r.Run(port)

}
