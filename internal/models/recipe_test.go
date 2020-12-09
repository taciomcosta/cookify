package models

import (
	"strings"
	"testing"

	"github.com/taciomcosta/cookify/pkg/recipepuppy"
)

var mockPuppyRecipe = []recipepuppy.Recipe{
	{
		Title:       "title",
		Href:        "href",
		Ingredients: "unsorted, ingredients, with, spaces",
		Thumbnail:   "thumbnail",
	},
}

var expectedRecipe = []Recipe{
	{
		Title:       "title",
		Ingredients: Ingredients{"ingredients", "spaces", "unsorted", "with"},
		Link:        "href",
		Gif:         "",
	},
}

func TestParseManyRecipes(t *testing.T) {
	results := ParseManyRecipes(mockPuppyRecipe)
	for i := range expectedRecipe {
		expectedIngredientsJoined := strings.Join(expectedRecipe[i].Ingredients, ",")
		resultIngredientsJoined := strings.Join(results[i].Ingredients, ",")
		if expectedIngredientsJoined != resultIngredientsJoined {
			t.Error("should parse ingredients")
		}

		if expectedRecipe[i].Title != results[i].Title {
			t.Error("should parse title")
		}

		if expectedRecipe[i].Link != results[i].Link {
			t.Error("should parse link")
		}

		if expectedRecipe[i].Gif != results[i].Gif {
			t.Error("should parse gif")
		}
	}
}
