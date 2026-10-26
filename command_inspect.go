package main

import (
	"errors"
	"fmt"
)

// commandInspect checks the caughtPokemon map for the Pokemon
// that inspect was called on, if caught displays stats about that pokemon
func commandInspect(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("entering the name of a Pokemon is required")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		fmt.Printf("you have not caught that pokemon\n")
		return nil

	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, spec := range pokemon.Types {
		fmt.Printf(" - %s\n", spec.Type.Name)
	}
	return nil

}
