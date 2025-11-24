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

func CExpl(cfg *Config, cache *pokecache.Cache, places []string) error {
	if len(places) == 0 {
		fmt.Println("No location given")
		return nil
	}
	url := "https://pokeapi.co/api/v2/location-area/" + places[0] + "/"

	fmt.Printf("Exploring %s...\n", places[0])
	data_place, err := helpers.RetrieveCache[Place](url, cache)
	if err != nil {
		return err
	}
	fmt.Printf("Found Pokemon:\n")
	for _, encounter := range data_place.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

/*
old CExpl entire region

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
*/
