package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"quiz-app-api/models"
)

func GetThemesOrLevels(c *gin.Context) {
	subCategoryID := c.Query("sub_category_id")
	if subCategoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sub_category_id query parameter is required"})
		return
	}

	fileData, err := os.ReadFile("data/category/themes_or_levels.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
		return
	}

	var response models.ThemesOrLevelsResponse
	err = json.Unmarshal(fileData, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
		return
	}

	subCategoryIDInt, err := strconv.Atoi(subCategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sub_category_id parameter"})
		return
	}

	var filteredThemesOrLevels []models.ThemeOrLevel
	for _, item := range response.Data {
		if item.SubCategoryID == subCategoryIDInt {
			filteredThemesOrLevels = append(filteredThemesOrLevels, item)
		}
	}

	if len(filteredThemesOrLevels) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No themes or levels found for the given sub_category_id"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": filteredThemesOrLevels})
}
