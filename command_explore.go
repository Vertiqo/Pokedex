package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) == 0 {
		return errors.New("entering the name of an area is required")
	}

	url := baseURL + args[0]
	encounters, err := cfg.client.FetchEncounters(url)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", args[0])
	fmt.Println("Found Pokemon:")
	for _, enc := range encounters.Encounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}
