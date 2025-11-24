package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gustavbagger/Pokedex/commands"
	pokecache "github.com/gustavbagger/Pokedex/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &commands.Config{
		Next:      "https://pokeapi.co/api/v2/location-area/",
		Previous:  "",
		Currently: "",
		Pokedex:   map[string]commands.Pokemon{},
	}
	interval := 5 * time.Second
	cache := pokecache.NewCache(interval)
	commandsMap := commands.Support(cfg, cache)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.Fields(scanner.Text())
		var command string
		var arg string
		switch len(input) {
		case 1:
			command = input[0]
		case 2:
			command = input[0]
			arg = input[1]
		default:
			fmt.Print("Unknown command\n")
			continue
		}

		if cli, ok := commandsMap[command]; ok {
			if err := cli.Callback(arg); err != nil {
				fmt.Printf("callback error: %v\n", err)
				return
			}

		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
