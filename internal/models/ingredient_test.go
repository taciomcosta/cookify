package models

import (
	"testing"
)

func TestParseIngredients(t *testing.T) {
	result, err := ParseIngredients("tomato, onion")

	if err != nil {
		t.Errorf("should parse array of ingredients (sorted)")
	}

	expected := Ingredients{"onion", "tomato"}
	for i := range expected {
		if expected[i] != result[i] {
			t.Errorf("should parse array of ingredients (sorted)")
		}
	}
}

func TestParseIngredientsError(t *testing.T) {
	_, err := ParseIngredients("tomato, onion, banana, pepper")

	if err == nil || err.Error() != "Please, provide at most 3 ingredients" {
		t.Error("should fail to parse more than 3 ingredients")
	}
}
