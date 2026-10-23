package main

import (
	"fmt"
)

// commandMap displays the next page of locations and updates pagination.

func commandMap(cfg *config, args []string) error {

	response, err := cfg.client.FetchLocations(*cfg.mapUrlNext)
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
