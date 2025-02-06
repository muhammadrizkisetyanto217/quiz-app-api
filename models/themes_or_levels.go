package models

type ThemeOrLevel struct {
	ID                      int     `json:"id"`
	Name                    string  `json:"name"`
	DescriptionShort        string  `json:"description_short"`
	DescriptionLong         string  `json:"description_long"`
	Status                  string  `json:"status"`
	GradeTotalUnit          int     `json:"grade_total_unit"`
	TotalUnit               int     `json:"total_unit"`
	CompletedQuizzesArticle int     `json:"completed_quizzes_article"`
	TotalQuizzesArticle     int     `json:"total_quizzes_article"`
	ImageURL                string  `json:"image_url"`
	CreatedAt               string  `json:"created_at"`
	UpdatedAt               *string `json:"updated_at"`
	DeletedAt               *string `json:"deleted_at"`
	SubCategoryID           int     `json:"sub_category_id"`
}

type ThemesOrLevelsResponse struct {
	Data []ThemeOrLevel `json:"data"`
}
