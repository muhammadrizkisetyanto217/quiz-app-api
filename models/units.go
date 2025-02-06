package models

type Unit struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	DescriptionShort   string  `json:"description_short"`
	DescriptionOverview string  `json:"description_overview"`
	Status             string  `json:"status"`
	GradeTotal         int     `json:"grade_total"`
	CompletedQuizzes   int     `json:"completed_quizzes"`
	TotalQuizzes       int     `json:"total_quizzes"`
	CompletedArticle   bool    `json:"completed_article"`
	IconOverviewURL    string  `json:"icon_overview_url"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          *string `json:"updated_at"`
	DeletedAt          *string `json:"deleted_at"`
	ThemesOrLevelID    int     `json:"themes_or_level_id"`
	CreatedBy          int     `json:"created_by"`
}

type UnitsResponse struct {
	Data []Unit `json:"data"`
}
