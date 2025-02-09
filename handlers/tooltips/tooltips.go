package tooltips

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"quiz-app-api/models"
)

// GetTooltips retrieves tooltips based on keyword query parameter
func GetTooltips(c *gin.Context) {
	keyword := c.Query("keyword")

	// Baca file JSON
	fileData, err := os.ReadFile("data/tooltips/tooltips.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
		return
	}

	// Parse JSON ke struct
	var response models.TooltipsResponse
	err = json.Unmarshal(fileData, &response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
		return
	}

	// Jika tidak ada keyword, kembalikan semua data
	if keyword == "" {
		c.JSON(http.StatusOK, gin.H{"data": response.Data})
		return
	}

	// Filter data berdasarkan keyword
	var filteredTooltips []models.Tooltip
	for _, item := range response.Data {
		if item.Keyword == keyword {
			filteredTooltips = append(filteredTooltips, item)
		}
	}

	// Jika tidak ada hasil setelah filter, kirim pesan bahwa tidak ditemukan
	if len(filteredTooltips) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tooltips found for the given keyword"})
		return
	}

	// Kirim hasil yang telah difilter
	c.JSON(http.StatusOK, gin.H{"data": filteredTooltips})
}
