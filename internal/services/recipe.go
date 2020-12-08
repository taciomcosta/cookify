package services

import (
	"errors"
	"strings"
	"sync"

	"github.com/taciomcosta/cookify/internal/models"
	"github.com/taciomcosta/cookify/pkg/giphy"
	"github.com/taciomcosta/cookify/pkg/recipepuppy"
)

type RecipeService struct{}

type FindByIngredientsResponse struct {
	Keywords models.Ingredients `json:"keywords"`
	Recipes  []models.Recipe    `json:"recipes"`
}

func (s RecipeService) FindByIngredients(ingredients string) (FindByIngredientsResponse, error) {
	err := checkServicesStatuses()
	if err != nil {
		return FindByIngredientsResponse{}, err
	}

	puppyRecipes, err := recipepuppy.FindRecipes(ingredients, "", 1)
	if err != nil {
		return FindByIngredientsResponse{}, err
	}
	recipes := parseManyRecipes(puppyRecipes)
	enhanceManyWithGifs(recipes)

	keywords, err := models.ParseIngredients(ingredients)
	if err != nil {
		return FindByIngredientsResponse{}, err
	}

	response := FindByIngredientsResponse{
		Recipes:  recipes,
		Keywords: keywords,
	}
	return response, nil
}

func checkServicesStatuses() error {
	_, err := recipepuppy.FindRecipes("", "", 1)
	if err != nil {
		return errors.New("Service unavailable: RecipePuppy")
	}

	_, err = giphy.Search("")
	if err != nil {
		return errors.New("Service unavailable: Giphy")
	}

	return nil
}

func parseManyRecipes(puppyRecipes []recipepuppy.PuppyRecipeDTO) []models.Recipe {
	recipes := make([]models.Recipe, 0)
	for _, puppyRecipe := range puppyRecipes {
		recipe := parseOneRecipe(puppyRecipe)
		recipes = append(recipes, recipe)
	}
	return recipes
}

func parseOneRecipe(puppyRecipe recipepuppy.PuppyRecipeDTO) models.Recipe {
	return models.Recipe{
		Title:       puppyRecipe.Title,
		Ingredients: strings.Split(puppyRecipe.Ingredients, ","),
		Link:        puppyRecipe.Href,
		Gif:         "",
	}
}

func enhanceManyWithGifs(recipes []models.Recipe) {
	var group sync.WaitGroup
	group.Add(len(recipes))
	for i := range recipes {
		go enhanceOneWithGif(&recipes[i], &group)
	}
	group.Wait()
}

func enhanceOneWithGif(recipe *models.Recipe, group *sync.WaitGroup) {
	recipe.Gif = findGif(recipe.Title)
	group.Done()
}

func findGif(title string) string {
	response, err := giphy.Search(title)
	if err != nil {
		return ""
	}

	if len(response.Data) == 0 {
		return ""
	}

	return response.Data[0].Images.Original.Url
}
