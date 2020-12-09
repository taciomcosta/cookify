package models

// Recipe represts a recipe in the domain model.
type Recipe struct {
	Title       string      `json:"title"`
	Ingredients Ingredients `json:"ingredients"`
	Link        string      `json:"link"`
	Gif         string      `json:"gif"`
}
