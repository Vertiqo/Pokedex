package main

import (
	"fmt"
)

// commandMapb displays the previous page of locations and updates pagination.
// If already on the first page, it prints a message instead.
func commandMapb(cfg *config, args []string) error {

	if cfg.mapUrlPrevious == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	response, err := cfg.client.FetchLocations(*cfg.mapUrlPrevious)
	if err != nil {
		return err
	}
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	cfg.mapUrlNext = response.Next
	cfg.mapUrlPrevious = response.Previous
	return nil

}
