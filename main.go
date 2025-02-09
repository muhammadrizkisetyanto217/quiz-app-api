package main

import (
	"quiz-app-api/handlers/category"
    "quiz-app-api/handlers/tooltips"

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


	r.GET("/subcategory", category.GetSubcategory)
	r.GET("/category", category.GetCategory)
	r.GET("/themes_or_levels", category.GetThemesOrLevels)
	r.GET("/units", category.GetUnits)
	r.GET("/tooltips", tooltips.GetTooltips)

	r.Run(":8080") // Server berjalan di http://localhost:8080
}
