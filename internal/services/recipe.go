package services

import (
	"errors"
	"sync"

	"github.com/taciomcosta/cookify/internal/models"
	"github.com/taciomcosta/cookify/pkg/giphy"
	"github.com/taciomcosta/cookify/pkg/recipepuppy"
)

// RecipeService makes application rules of recipes
// available for external adapters to be plugged in (http, amqp, grpc)
type RecipeService struct{}

// FindByIngredientsResponse represents the format
// of data returned by FindByIngredients.
type FindByIngredientsResponse struct {
	Keywords models.Ingredients `json:"keywords"`
	Recipes  []models.Recipe    `json:"recipes"`
}

// FindByIngredients finds recipes and its gifs
// by consuming external services.
func (s RecipeService) FindByIngredients(ingredients string) (FindByIngredientsResponse, error) {
	err := checkServicesStatuses()
	if err != nil {
		return FindByIngredientsResponse{}, err
	}

	puppyRecipes, err := recipepuppy.FindRecipes(ingredients, "", 1)
	if err != nil {
		return FindByIngredientsResponse{}, err
	}
	recipes := models.ParseManyRecipes(puppyRecipes)
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

	return response.Data[0].Images.Original.URL
}
