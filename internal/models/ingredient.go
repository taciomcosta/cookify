package models

import (
	"errors"
	"strings"
)

type Ingredients []string
type Ingredient string

const maxIngredientsAllowd = 3

func ParseIngredients(ingredients string) (Ingredients, error) {
	strings := strings.Split(ingredients, ",")
	if len(strings) > maxIngredientsAllowd {
		return nil, errors.New("Please, provide at most 3 ingredients")
	}
	return strings, nil
}
