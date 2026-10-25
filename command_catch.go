package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("entering the name of a Pokemon is required")
	}

	url := baseURL + "pokemon/" + args[0]
	pokemon, err := cfg.client.FetchPokemon(url)
	if err != nil {
		return err
	}
	// check for http code
	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	attempt := rand.Intn(pokemon.Exp)
	fmt.Printf("chances: %d to %d\n", attempt, pokemon.Exp)
	if attempt < 50 {
		fmt.Printf("%s was caught!\n", args[0])
		cfg.caughtPokemon[pokemon.Name] = pokemon
		return nil
	}
	fmt.Printf("%s escaped!\n", args[0])
	return nil

}
