package main

import (
	"quiz-app-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/subcategory", handlers.GetSubcategory)

	r.Run(":8080") // Server berjalan di http://localhost:8080
}
