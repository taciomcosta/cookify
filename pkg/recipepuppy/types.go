package recipepuppy

// RecipeResponse represents main response returned by RecipePuppy's API.
type RecipeResponse struct {
	Title   string   `json:"title"`
	Version float32  `json:"version"`
	Href    string   `json:"href"`
	Results []Recipe `json:"results"`
}

// Recipe represents a recipe returned by RecipePuppy's API.
type Recipe struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail   string `json:"thumbnail"`
}
