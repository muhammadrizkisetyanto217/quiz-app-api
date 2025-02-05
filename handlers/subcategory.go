package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"quiz-app-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSubcategory(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
		return
	}

	fileData, err := os.ReadFile("data/subcategory.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
		return
	}

	var response models.Response
	err = json.Unmarshal(fileData, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
		return
	}

	for _, category := range response.Data {
		if category.ID == idInt {
			c.JSON(http.StatusOK, category)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
}
