package recipepuppy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// TODO: export to env var
var url = "http://www.recipepuppy.com/api/"
var client = &http.Client{}

func FindRecipes(ingredients string, query string, page int) ([]PuppyRecipeDTO, error) {
	request, err := newRequest(ingredients, query, page)
	if err != nil {
		return []PuppyRecipeDTO{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return []PuppyRecipeDTO{}, err
	}

	return parsePuppyRecipes(response)
}

func newRequest(ingredients string, query string, page int) (*http.Request, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New("error on creating request")
	}

	queryString := request.URL.Query()
	queryString.Add("i", ingredients)
	queryString.Add("q", query)
	queryString.Add("p", strconv.Itoa(page))
	request.URL.RawQuery = queryString.Encode()

	return request, nil
}

func parsePuppyRecipes(response *http.Response) ([]PuppyRecipeDTO, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []PuppyRecipeDTO{}, errors.New("failed parsing Puppy Recipe response")
	}

	var parsedResponse PuppyRecipeResponse
	err = json.Unmarshal(bytes, &parsedResponse)
	if err != nil {
		return []PuppyRecipeDTO{}, errors.New("failed parsing Puppy Recipe response")
	}

	return parsedResponse.Results, err
}
