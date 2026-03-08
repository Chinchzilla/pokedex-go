package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/Chinchzilla/pokedex-go/internal/pokedexapi"
)

const cacheInterval = 5 * time.Second

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		pokedexAPIClient: pokedexapi.GetPokedexAPIClient(cacheInterval),
		nextUrl:          nil,
		prevUrl:          nil,
	}
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

		input := scanner.Text()
		cliCommands := getCommands()

		command, ok := cliCommands[input]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback(cfg)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}

}
