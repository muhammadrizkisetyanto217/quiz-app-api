package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Category struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ArticlesCount int    `json:"articles_count"`
}

type Response struct {
	Data []Category `json:"data"`
}

func getSubcategory(w http.ResponseWriter, r *http.Request) {
	// Ambil parameter query id dari URL
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id query parameter is required", http.StatusBadRequest)
		return
	}

	// Baca file JSON
	fileData, err := ioutil.ReadFile("subcategory.json")
	if err != nil {
		http.Error(w, "Error reading the JSON file", http.StatusInternalServerError)
		return
	}

	// Parse JSON ke struct
	var response Response
	err = json.Unmarshal(fileData, &response)
	if err != nil {
		http.Error(w, "Error parsing JSON", http.StatusInternalServerError)
		return
	}

	// Convert id ke integer
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid id parameter", http.StatusBadRequest)
		return
	}

	// Cari category berdasarkan id
	for _, category := range response.Data {
		if category.ID == idInt {
			// Jika ditemukan, kirimkan response JSON
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
			return
		}
	}

	// Jika tidak ditemukan
	http.Error(w, "Category not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/subcategory", getSubcategory)

	port := ":8080"
	fmt.Println("Server running at http://localhost" + port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}