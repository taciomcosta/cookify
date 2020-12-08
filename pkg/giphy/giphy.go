package giphy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// TODO: export to env var
var url = "https://api.giphy.com/v1/gifs/search"
var apiKey = "pPiMNFkdnBt4wGmBiJ9YCryAw3lHJk98"
var client = &http.Client{}

func Search(query string) (*GiphySearchResponse, error) {
	request, err := newRequest(query)
	if err != nil {
		return &GiphySearchResponse{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return &GiphySearchResponse{}, err
	}

	return parseResponse(response)
}

func newRequest(query string) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("error on creating request for Giphy")
	}

	queryString := request.URL.Query()
	queryString.Add("q", query)
	queryString.Add("api_key", apiKey)
	queryString.Add("limit", "1")
	request.URL.RawQuery = queryString.Encode()

	return request, nil
}

func parseResponse(response *http.Response) (*GiphySearchResponse, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &GiphySearchResponse{}, errors.New("failed parsing Giphy response")
	}

	var parsedResponse GiphySearchResponse
	err = json.Unmarshal(bytes, &parsedResponse)
	if err != nil {
		return &GiphySearchResponse{}, errors.New("failed parsing Giphy response")
	}

	return &parsedResponse, err
}
