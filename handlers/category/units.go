package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"quiz-app-api/models"
)

func GetUnits(c *gin.Context) {
	subCategoryID := c.Query("themes_or_levels")
	if subCategoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "themes_or_levels query parameter is required"})
		return
	}

	fileData, err := os.ReadFile("data/category/units.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
		return
	}

	var response models.UnitsResponse
	err = json.Unmarshal(fileData, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
		return
	}

	subCategoryIDInt, err := strconv.Atoi(subCategoryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid themes_or_levels parameter"})
		return
	}

	var filteredThemesOrLevels []models.Unit
	for _, item := range response.Data {
		if item.ThemesOrLevelID == subCategoryIDInt {
			filteredThemesOrLevels = append(filteredThemesOrLevels, item)
		}
	}

	if len(filteredThemesOrLevels) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No themes or levels found for the given themes_or_levels"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": filteredThemesOrLevels})
}
