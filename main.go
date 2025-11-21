package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavbagger/Pokedex/helpers"
)

func main() {
	Supported := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    CExit,
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := helpers.CleanInput(text)
		if len(cleaned) == 0 {
			fmt.Println("Your command was: ")
		} else {
			fmt.Printf(
				"Your command was: %s\n",
				cleaned[0],
			)
		}
	}
}
