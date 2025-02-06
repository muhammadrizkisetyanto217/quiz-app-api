package main

import (
	"quiz-app-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight request
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

	r.GET("/subcategory", handlers.GetSubcategory)
	r.GET("/category", handlers.GetCategory)
	r.GET("/themes_or_levels", handlers.GetThemesOrLevels)

	r.Run(":8080") // Server berjalan di http://localhost:8080
}
