package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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

		command.callback()
	}

}
