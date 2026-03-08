package main

import (
	"fmt"
	"os"

	"github.com/Chinchzilla/pokedex-go/internal/pokedexapi"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, args ...string) error {
	cliCommands := getCommands()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")
	for _, command := range cliCommands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}

func commandMap(cfg *Config, args ...string) error {
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

func commandMapBackward(cfg *Config, args ...string) error {
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

func commandExplore(cfg *Config, args ...string) error {
	location := args[0]
	fmt.Printf("Exploring %s...\n", location)

	res, err := pokedexapi.ExploreLocation(cfg.pokedexAPIClient, location)
	if err != nil {
		return err
	}

	for _, pokemon := range res.PokemonEncounters {
		fmt.Printf("\t- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
