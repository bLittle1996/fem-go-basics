package api_test

import (
	"front-end-masters/go-basics/working-with-api/api"
	"testing"
)

func TestGetPokemonShouldErrorWhenProvidedWithAnEmptyInput(t *testing.T) {
	_, err := api.GetPokemon("")

	if err == nil {
		t.Error("Expected error when given empty input")
	}
}

func TestGetPokemonShouldReturnAPokemonWhenProvidedValidInput(t *testing.T) {
	type testCase struct {
		input        string
		expectedName string
		expectedId   int
	}

	cases := []testCase{
		{"bulbasaur", "bulbasaur", 1},
		{"pikachu", "pikachu", 25},
		{"1", "bulbasaur", 1},
		{"25", "pikachu", 25},
	}

	for _, testCase := range cases {
		pokemon, err := api.GetPokemon(testCase.input)

		if err != nil {
			t.Errorf("Unexpected error for input %s: %v", testCase.input, err)
		} else if pokemon.Id != testCase.expectedId {
			t.Errorf("Id mismatch: got %d, expected %d", pokemon.Id, testCase.expectedId)
		} else if pokemon.Name != testCase.expectedName {
			t.Errorf("Name mismatch: got %s, expected %s", pokemon.Name, testCase.expectedName)
		}
	}
}

func TestGetPokemonShouldErrorIfGivenInvalidInput(t *testing.T) {
	_, err := api.GetPokemon("notarealdigimon")

	if err == nil {
		t.Errorf("Expected error when given invalid input")
	}
}
