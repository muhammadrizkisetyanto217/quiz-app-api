package models

// Tooltip represents the structure of tooltip data
type Tooltip struct {
	ID           uint   `json:"id"`
	Keyword      string `json:"keyword"`
	TooltipShort string `json:"tooltip_short"`
	TooltipLong  string `json:"tooltip_long"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type TooltipsResponse struct {
	Data []Tooltip `json:"data"`
}