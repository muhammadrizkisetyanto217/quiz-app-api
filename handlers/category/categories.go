package handlers

import (
    "encoding/json"
    "net/http"
    "os"
    "strconv"

    "quiz-app-api/models"
    "github.com/gin-gonic/gin"
)

func GetCategory(c *gin.Context) {
    id := c.Query("id") // Ambil query parameter 'id'

    // Jika tidak ada parameter 'id', kembalikan seluruh data kategori
    if id == "" {
        fileData, err := os.ReadFile("data/category/categories.json")
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
            return
        }

        var response models.CategoryResponse
        err = json.Unmarshal(fileData, &response)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
            return
        }

        c.JSON(http.StatusOK, response) // Kembalikan seluruh data kategori
        return
    }

    // Jika ada parameter 'id', cari kategori dengan ID tersebut
    fileData, err := os.ReadFile("data/category.json")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading the JSON file"})
        return
    }

    var response models.CategoryResponse
    err = json.Unmarshal(fileData, &response)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
        return
    }

    idInt, err := strconv.Atoi(id) // Mengonversi id menjadi integer
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id parameter"})
        return
    }

    for _, category := range response.Data {
        if category.ID == idInt {
            c.JSON(http.StatusOK, category) // Kembalikan kategori yang ditemukan
            return
        }
    }

    c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"}) // Jika kategori tidak ditemukan
}
