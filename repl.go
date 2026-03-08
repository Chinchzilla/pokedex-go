package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Chinchzilla/pokedex-go/internal/pokedexapi"
)

type Config struct {
	pokedexAPIClient *pokedexapi.PokeClient
	nextUrl          *string
	prevUrl          *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display help information",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Map location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Map location areas backwards",
			callback:    commandMapBackward,
		},
		"explore": {
			name:        "explore",
			description: "Explore a given area",
			callback:    commandExplore,
		},
	}
}

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	// Initialize empty string for nextUrl as a first value
	// The client should handle it as the first request and initialize
	// the corret nextUrl field with the URL returned by the API.
	var emptyString string
	cfg.nextUrl = &emptyString

	for {
		fmt.Print("Pokedex > ")

		if !scanner.Scan() {
			break
		}

		input := cleanInput(scanner.Text())
		cliCommands := getCommands()

		command, ok := cliCommands[input[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		var args []string
		if len(input) > 1 {
			args = input[1:]
		}

		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func cleanInput(text string) []string {
	trim := strings.TrimSpace(text)
	lowerCase := strings.ToLower(trim)
	split := strings.Fields(lowerCase)
	return split
}
