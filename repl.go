package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	var result []string

	loweredText := strings.ToLower(text)
	result = strings.Fields(loweredText)

	return result

}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		query := scanner.Text()
		clean := cleanInput(query)
		if len(clean) == 0 {
			continue
		}
		value, ok := cfg.availableCommands[clean[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		// value.callback(cfg) actually invokes the matched command (e.g. commandHelp).
		// It returns nil on success, or a non-nil error if something went wrong.
		if err := value.callback(cfg); err != nil {
			fmt.Println("Error:", err)
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
