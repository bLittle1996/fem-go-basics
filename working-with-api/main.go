package main

import (
	"fmt"
	"front-end-masters/go-basics/working-with-api/api"
)

func main() {
	pikachu, err := api.GetPokemon("pikachu")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pikachu)
}
