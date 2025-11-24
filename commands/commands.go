package commands

import (
	"fmt"
	"os"

	"github.com/gustavbagger/Pokedex/helpers"
	pokecache "github.com/gustavbagger/Pokedex/internal"
)

func CExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CHelp(cfg *Config, cache *pokecache.Cache) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for _, value := range Support(cfg, cache) {
		fmt.Printf("\n%v: %v", value.Name, value.Description)
	}
	fmt.Println("")
	return nil
}

func CMap(cfg *Config, cache *pokecache.Cache) error {
	return Maps(cfg.Next, cfg, cache)
}

func CMapb(cfg *Config, cache *pokecache.Cache) error {
	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	return Maps(cfg.Previous, cfg, cache)
}

func CExpl(cfg *Config, cache *pokecache.Cache) error {

	data, err := helpers.RetrieveCache[LocationArea](cfg.Currently, cache)
	if err != nil {
		return err
	}
	pokemon := map[string]bool{}
	for _, area := range data.Results {
		data_area, err := helpers.RetrieveCache[Place](area.URL, cache)
		if err != nil {
			return err
		}
		for _, encounter := range data_area.PokemonEncounters {
			pokemon[encounter.Pokemon.Name] = true
		}
	}
	for key := range pokemon {
		fmt.Println(key)
	}

	return nil
}
