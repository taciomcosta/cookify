package recipepuppy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/taciomcosta/cookify/internal/config"
)

var url string
var client *http.Client

func init() {
	url = config.GetString("RECIPEPUPPY_URL")
	client = &http.Client{}
}

// FindRecipes wraps recipes endpoint of RecipePuppy's API.
func FindRecipes(ingredients string, query string, page int) ([]Recipe, error) {
	request, err := newRequest(ingredients, query, page)
	if err != nil {
		return []Recipe{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		return []Recipe{}, err
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

func parsePuppyRecipes(response *http.Response) ([]Recipe, error) {
	bytes, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	if err != nil {
		return []Recipe{}, errors.New("failed parsing Puppy Recipe response")
	}

	var parsedResponse RecipeResponse
	err = json.Unmarshal(bytes, &parsedResponse)
	if err != nil {
		return []Recipe{}, errors.New("failed parsing Puppy Recipe response")
	}

	return parsedResponse.Results, err
}
