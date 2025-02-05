package models

type Subcategory struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	ArticlesCount int    `json:"articles_count"`
}

type Response struct {
	Data []Subcategory `json:"data"`
}
