package models

type UpdateNews struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	SubTitle    string `json:"sub_title"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type SubCategory struct {
	ID                      int          `json:"id"`
	Name                    string       `json:"name"`
	DescriptionLong         string       `json:"description_long"`
	Status                  string       `json:"status"`
	GradeTotalThemesOrLevel int          `json:"grade_total_themes_or_level"`
	TotalThemesOrLevel      int          `json:"total_themes_or_level"`
	CompletedThemesOrLevels int          `json:"completed_themes_or_levels"`
	UpdateNews              []UpdateNews `json:"update_news"`
	ImageURL                string       `json:"image_url"`
	CreatedAt               string       `json:"created_at"`
	UpdatedAt               *string      `json:"updated_at"` // Bisa null
	CategoryID              int          `json:"category_id"`
}

type Response struct {
	Data []SubCategory `json:"data"`
}
