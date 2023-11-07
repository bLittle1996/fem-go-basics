package api

import (
	"encoding/json"
	"fmt"
	"front-end-masters/go-basics/working-with-api/pokemon"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2"
const pokemonEndpoint = baseUrl + "/pokemon/%s"

func GetPokemon(idOrName string) (*pokemon.Pokemon, error) {
	if idOrName == "" {
		return nil, fmt.Errorf("idOrName must not be empty")
	}

	resp, err := http.Get(fmt.Sprintf(pokemonEndpoint, idOrName))

	if err != nil {
		return nil, err
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error %d: unable to request data for %s", resp.StatusCode, idOrName)
	}

	var body []byte
	body, err = io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	pokemon := &pokemon.Pokemon{}

	err = json.Unmarshal(body, pokemon)

	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
