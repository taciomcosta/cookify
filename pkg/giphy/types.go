package giphy

type GiphySearchResponse struct {
	Data []GiphyDTO `json:"data"`
}

type GiphyDTO struct {
	Images GiphyImage `json:"images"`
}

type GiphyImage struct {
	Original GiphyImageOriginal `json:"original"`
}

type GiphyImageOriginal struct {
	Url string `json:"url"`
}
