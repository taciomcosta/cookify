package recipepuppy

type PuppyRecipeDTO struct {
	Title       string `json:"title"`
	Href        string `json:"href"`
	Ingredients string `json:"ingredients"`
	Thumbnail   string `json:"thumbnail"`
}

type PuppyRecipeResponse struct {
	Title   string           `json:"title"`
	Version float32          `json:"version"`
	Href    string           `json:"href"`
	Results []PuppyRecipeDTO `json:"results"`
}
