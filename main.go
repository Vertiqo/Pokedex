package main

import (
	"time"

	"github.com/vertiqo/pokedex/internal/pokeapi"
)

// Pokeapi base url:

var baseURL = "https://pokeapi.co/api/v2/location-area/"

//struct to hold and hand over shared state
// Next and Previous are pointers rather than plain strings because the
// PokeAPI returns JSON `null` for these fields when there is no next or
// previous page. A plain string can't represent "no value", it would just
// be an empty string "", which is ambiguous (is it empty because there's
// no more pages, or because the API actually returned an empty string?).
// A *string can be nil, which cleanly and unambiguously represents
// JSON's null.

type config struct {
	availableCommands map[string]cliCommand
	mapUrlPrevious    *string
	mapUrlNext        *string
	client            pokeapi.Client
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// builds the map of commands available to the user later to be
// stored in the config struct

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays areas of the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous areas",
			callback:    commandMapb,
		},
	}
}

func main() {

	// the config struct gets populated and put in the variable cfg
	// this cfg instance gets passed everywhere
	// mapUrlNext has to lead somewhere, in this case the first page fits perfectly
	// a new client

	cfg := &config{
		availableCommands: getCommands(),
		mapUrlNext:        &baseURL,
		client:            pokeapi.NewClient(5 * time.Second),
	}
	startRepl(cfg)

}
