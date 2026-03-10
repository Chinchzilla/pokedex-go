package main

import (
	"fmt"
	"math/rand"
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
	if len(args) == 0 {
		return fmt.Errorf("No location provided")
	}
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

func commandCatch(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon name provided")
	}
	pokemon_name := args[0]
	_, caught := cfg.caughtPokemon[pokemon_name]
	if caught {
		fmt.Printf("%s is already caught!\n", pokemon_name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon_name)
	pokemon, err := pokedexapi.GetPokemon(cfg.pokedexAPIClient, pokemon_name)
	if err != nil {
		fmt.Printf("Error getting pokemon: %s\n", err)
		return err
	}

	chance := rand.Intn(500)
	if chance > pokemon.BaseExperience {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("No pokemon name provided")
	}

	pokemon_name := args[0]
	pokemon, caught := cfg.caughtPokemon[pokemon_name]
	if !caught {
		return fmt.Errorf("%s is not caught", pokemon.Name)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, pokeType := range pokemon.Types {
		fmt.Printf("\t- %s\n", pokeType.Type.Name)
	}

	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("No pokemon caught")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}
