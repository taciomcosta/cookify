package giphy

// SearchResponse represents main format
// returned by Giphy API
type SearchResponse struct {
	Data []Gif `json:"data"`
}

// Gif represents gif metadata
// returned by Giphy API
type Gif struct {
	Images Image `json:"images"`
}

// Image represents image metadata
// returned by Giphy API
type Image struct {
	Original ImageOriginal `json:"original"`
}

// ImageOriginal represents original image metadata
// returned by Giphy API
type ImageOriginal struct {
	URL string `json:"url"`
}
