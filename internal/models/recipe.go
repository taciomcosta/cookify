package models

import (
	"sort"
	"strings"

	"github.com/taciomcosta/cookify/pkg/recipepuppy"
)

// Recipe represts a recipe in the domain model.
type Recipe struct {
	Title       string      `json:"title"`
	Ingredients Ingredients `json:"ingredients"`
	Link        string      `json:"link"`
	Gif         string      `json:"gif"`
}

// ParseManyRecipes parses recipes from RecipePuppy into our format.
func ParseManyRecipes(puppyRecipes []recipepuppy.Recipe) []Recipe {
	recipes := make([]Recipe, 0)
	for _, puppyRecipe := range puppyRecipes {
		recipe := parseOneRecipe(puppyRecipe)
		recipes = append(recipes, recipe)
	}
	return recipes
}

func parseOneRecipe(puppyRecipe recipepuppy.Recipe) Recipe {
	return Recipe{
		Title:       puppyRecipe.Title,
		Ingredients: parsePuppyIngredients(puppyRecipe.Ingredients),
		Link:        puppyRecipe.Href,
		Gif:         "",
	}
}

func parsePuppyIngredients(ingredients string) Ingredients {
	parsed := strings.Split(ingredients, ",")

	for i := range parsed {
		parsed[i] = strings.TrimSpace(parsed[i])
	}

	sort.Strings(parsed)
	return parsed
}
