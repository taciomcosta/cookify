package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taciomcosta/cookify/internal/models"

	"github.com/taciomcosta/cookify/internal/services"
)

var recipeService models.RecipeService

func init() {
	recipeService = services.NewRecipeService()
}

func FindRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, err := recipeService.FindByIngredients("")

	bytes, err := json.Marshal(recipes)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(bytes)
}
