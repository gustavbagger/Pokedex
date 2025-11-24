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
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}
	interval := 5 * time.Second
	cache := pokecache.NewCache(interval)
	commandsMap := commands.Support(cfg, cache)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := strings.Fields(scanner.Text())

		command := input[0]
		args := input[1:]

		if cli, ok := commandsMap[command]; ok {
			if err := cli.Callback(args); err != nil {
				fmt.Printf("callback error: %v\n", err)
				return
			}

		} else {
			fmt.Print("Unknown command\n")
		}
	}
}
