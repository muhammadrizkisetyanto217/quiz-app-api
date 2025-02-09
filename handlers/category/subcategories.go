package category

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"quiz-app-api/models"

	"github.com/gin-gonic/gin"
)

func GetSubcategory(c *gin.Context) {
	categoryID := c.Query("category_id") // Ambil parameter 'category_id' dari query string
	if categoryID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category_id query parameter is required"})
		return
	}

	// Membaca file category.json
	fileData, err := os.ReadFile("data/category/subcategories.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
		return
	}

	//* Cara 1
	// Membaca data ke dalam struct CategoryResponse
	// var response models.Response
	// err = json.Unmarshal(fileData, &response)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
	// 	return
	// }

	//* Cara 2
	response := &models.Response{} // Langsung sebagai pointer
	err = json.Unmarshal(fileData, response)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
		return
	}


	categoryIDInt, err := strconv.Atoi(categoryID) // Mengonversi 'category_id' menjadi integer
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category_id parameter"})
		return
	}

	// Mencari subkategori berdasarkan category_id
	// Filter berdasarkan category_id: Jika CategoryID dari subkategori sama dengan categoryIDInt, subkategori tersebut ditambahkan ke slice subcategories.
	var subcategories []models.SubCategory
	for _, subcategory := range response.Data {
		if subcategory.CategoryID == categoryIDInt {
			subcategories = append(subcategories, subcategory)
		}
	}

	// Jika tidak ada subkategori yang ditemukan
	if len(subcategories) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No subcategories found for the given category_id"})
		return
	}

	// Mengembalikan daftar subkategori yang sesuai dengan category_id
	c.JSON(http.StatusOK, subcategories)
}