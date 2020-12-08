package models

import (
	"fmt"
	"sort"
	"strings"
)

type Ingredients []string
type Ingredient string

const maxIngredientsAllowed = 3

func ParseIngredients(ingredients string) (Ingredients, error) {
	parsed := strings.Split(ingredients, ",")

	if len(parsed) > maxIngredientsAllowed {
		return nil, fmt.Errorf("Please, provide at most %d ingredients", maxIngredientsAllowed)
	}

	for i := range parsed {
		parsed[i] = strings.TrimSpace(parsed[i])
	}

	sort.Strings(parsed)
	return parsed, nil
}
