package main

import (
	"bufio"
	"fmt"
	"os"
	"time"


	"github.com/gustavbagger/Pokedex/commands"
	"github.com/gustavbagger/Pokedex/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	cfg := &commands.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: "",
	}
	cache := pokecache.NewCache(5*time.Second)

	commandsMap := commands.Support(cfg,cache)
	
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		command := scanner.Text()
		cache.ReapLoop()
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
