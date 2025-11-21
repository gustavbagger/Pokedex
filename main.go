package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gustavbagger/Pokedex/commands"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var Supported = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commands.CExit,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		if cli, ok := Supported[command]; ok {
			err := cli.callback
			if err != nil {
				fmt.Errorf("callback error: %v", err)
				return
			}
		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
