package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavbagger/Pokedex/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &commands.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}

	commandsMap := commands.Support(cfg)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()

		if cli, ok := commandsMap[command]; ok {
			if err := cli.Callback(); err != nil {
				fmt.Printf("callback error: %v\n", err)
				return
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
