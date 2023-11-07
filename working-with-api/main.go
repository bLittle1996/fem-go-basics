package main

import (
	"fmt"
	"front-end-masters/go-basics/working-with-api/api"
	"sync"
)

func main() {
	pokemans := []string{"pikachu", "raichu", "bulbasaur", "squirtle", "mew", "eevee"}
	var wg sync.WaitGroup

	for _, pokemonName := range pokemans {
		wg.Add(1) // Add one thing to the wait group for each pokemon

		// What if the function we want to call doesn't take a wait group? I guess we could make a wrapper around it?
		// Perhaps do a IIFE? Note we have to pass the variable in as a param otherwise each go routines variable would change to be the same one
		go func(pokeName string) {
			printPokemonDetails(pokeName)
			wg.Done()
		}(pokemonName)
	}

	fmt.Println("Waiting for pokemons to be summoned from the api")
	wg.Wait() // Waits for counter to hit zero
	fmt.Println("All calls finished!")
}

func printPokemonDetails(pokemon string) {
	pikachu, err := api.GetPokemon(pokemon)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(pikachu)
}
