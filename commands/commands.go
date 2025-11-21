package commands

import (
	"fmt"
	"os"
)

func CExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for _, value := range Support(cfg) {
		fmt.Printf("%v: %v\n", value.Name, value.Description)
	}
	return nil
}

func CMap(cfg *Config) error {
	return Maps(cfg.Next, cfg)
}

func CMapb(cfg *Config) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return Maps(cfg.Previous, cfg)
}
