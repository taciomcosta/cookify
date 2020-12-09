package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/taciomcosta/cookify/internal/services"
)

var recipeService = services.RecipeService{}

// FindRecipes is handler for finding recipes along with gifs
func FindRecipes(writer http.ResponseWriter, request *http.Request) {
	ingredients := request.URL.Query().Get("i")
	response, err := recipeService.FindByIngredients(ingredients)
	if err != nil {
		respondWithError(writer, err.Error(), http.StatusBadRequest)
		return
	}
	respondSuccess(writer, response)
}

func respondSuccess(writer http.ResponseWriter, value interface{}) {
	bytes, err := json.Marshal(value)
	if err != nil {
		respondWithError(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Write(bytes)
}

func respondWithError(writer http.ResponseWriter, message string, code int) {
	errorMessage := fmt.Sprintf(`{"error": "%s"}`, message)
	bytes := []byte(errorMessage)
	writer.WriteHeader(code)
	writer.Write(bytes)
}
