package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func cleanInput(text string) []string {
	trim := strings.TrimSpace(text)
	lowerCase := strings.ToLower(trim)
	split := strings.Fields(lowerCase)
	return split
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	cliCommands := getCommands()
	fmt.Println("Available commands:")
	for _, command := range cliCommands {
		fmt.Printf("%s - %s\n", command.name, command.description)
	}
	return nil
}
