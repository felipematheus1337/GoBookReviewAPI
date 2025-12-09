package main

import (
	"fmt"
	"os"

	"github.com/felipematheus1337/GoBookReviewAPI/config"

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

	_ = db

	r.Run(port)

}
