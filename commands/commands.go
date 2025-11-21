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
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, value := range Support(cfg) {
		fmt.Printf("\n%v: %v", value.Name, value.Description)
	}
	fmt.Println("")
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
