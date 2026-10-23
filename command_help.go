package main

import (
	"fmt"
)

func commandHelp(cfg *config, args []string) error {

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage:\n\n")
	// cfg.availableCommands is a shared state
	for _, cmd := range cfg.availableCommands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)

	}

	return nil
}
