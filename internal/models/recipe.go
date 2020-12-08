package models

type Recipe struct {
	Title       string      `json:"title"`
	Ingredients Ingredients `json:"ingredients"`
	Link        string      `json:"link"`
	Gif         string      `json:"gif"`
}
