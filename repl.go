package main

import (
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
	callback    func(*Config) error
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
	}
}

func cleanInput(text string) []string {
	trim := strings.TrimSpace(text)
	lowerCase := strings.ToLower(trim)
	split := strings.Fields(lowerCase)
	return split
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	cliCommands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	fmt.Println("Mapping location areas...")
	if cfg.nextUrl == nil {
		return fmt.Errorf("No more location areas found.")
	}
	res, err := pokedexapi.GetLocationArea(*cfg.nextUrl, cfg.pokedexAPIClient)
	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	cfg.nextUrl = res.Next
	cfg.prevUrl = res.Previous

	return nil
}

func commandMapBackward(cfg *Config) error {
	fmt.Println("Mapping location areas backwards...")
	if cfg.prevUrl == nil {
		return fmt.Errorf("No more location areas found.")
	}
	res, err := pokedexapi.GetLocationArea(*cfg.prevUrl, cfg.pokedexAPIClient)
	if err != nil {
		return err
	}

	for _, location := range res.Results {
		fmt.Println(location.Name)
	}

	cfg.nextUrl = res.Next
	cfg.prevUrl = res.Previous

	return nil
}
