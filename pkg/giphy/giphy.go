package giphy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/taciomcosta/cookify/internal/config"
)

var url string
var apiKey string
var client *http.Client

func init() {
	url = config.GetString("GLIPHY_URL")
	apiKey = config.GetString("GLIPHY_APIKEY")
	client = &http.Client{}
}

// Search wraps search endpoint of Giphy API
func Search(query string) (*SearchResponse, error) {
	request, err := newRequest(query)
	if err != nil {
		return &SearchResponse{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return &SearchResponse{}, err
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

func parseResponse(response *http.Response) (*SearchResponse, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return &SearchResponse{}, errors.New("failed parsing Giphy response")
	}

	var parsedResponse SearchResponse
	err = json.Unmarshal(bytes, &parsedResponse)
	if err != nil {
		return &SearchResponse{}, errors.New("failed parsing Giphy response")
	}

	return &parsedResponse, err
}
