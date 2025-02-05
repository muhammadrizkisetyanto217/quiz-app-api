package models

type Category struct {
	ID                int     `json:"id"`
	Name              string  `json:"name"`
	DescriptionShort  string  `json:"description_short"`
	DescriptionLong   string  `json:"description_long"`
	TotalSubcategories int    `json:"total_subcategories"`
	ImageURL          string  `json:"image_url"`
	Status            string  `json:"status"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         *string `json:"updated_at"` // Bisa null
}

type CategoryResponse struct {
	Data []Category `json:"data"`
}
