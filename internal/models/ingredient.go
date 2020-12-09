package models

import (
	"fmt"
	"sort"
	"strings"
)

// Ingredients represents many ingredients in the domain model.
type Ingredients []string

const maxIngredientsAllowed = 3

// ParseIngredients parses a comma separeted list of strings
// into a sorted slice of ingredients (Ingredients).
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
