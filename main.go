package main

import (
	"time"

	"github.com/vertiqo/pokedex/internal/pokeapi"
)

// Pokeapi base url:

var baseURL = "https://pokeapi.co/api/v2/"

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
	caughtPokemon     map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
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
		"explore": {
			name:        "explore",
			description: "Displays Pokemon local to that area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Lists the stats of a Pokemon if its in the collection",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lists all the caught Pokemon",
			callback:    commandPokedex,
		},
	}
}

func main() {

	// the config struct gets populated and put in the variable cfg
	// this cfg instance gets passed everywhere
	// mapUrlNext has to lead somewhere, in this case the first page fits perfectly
	// a new client
	startingURL := baseURL + "location-area/"

	cfg := &config{
		availableCommands: getCommands(),
		mapUrlNext:        &startingURL,
		client:            pokeapi.NewClient(5 * time.Second),
		caughtPokemon:     make(map[string]pokeapi.Pokemon),
	}
	startRepl(cfg)

}
