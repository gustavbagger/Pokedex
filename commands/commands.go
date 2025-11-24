package commands

import (
	"fmt"
	"math/rand"
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

func CExpl(cfg *Config, cache *pokecache.Cache, place string) error {

	url := "https://pokeapi.co/api/v2/location-area/" + place + "/"

	fmt.Printf("Exploring %s...\n", place)
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

func CCat(cfg *Config, cache *pokecache.Cache, pokemon string) error {
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon + "/"
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if _, ok := cfg.Pokedex[pokemon]; ok {
		fmt.Printf("%s is already in your Pokedex.\n", pokemon)
		return nil
	}
	data_pokemon, err := helpers.RetrieveCache[Pokemon](url, cache)
	if err != nil {
		return err
	}
	probability := 20.0 / float32(data_pokemon.BaseExperience)
	// fmt.Println(probability)
	if rand.Float32() < float32(probability) {
		cfg.Pokedex[pokemon] = data_pokemon
		fmt.Printf("%s was caught and added to your Pokedex!\n", pokemon)
	} else {
		fmt.Printf("%s ran away...\n", pokemon)
	}
	return nil
}

func CInsp(cfg *Config, cache *pokecache.Cache, pokemon string) error {
	if _, ok := cfg.Pokedex[pokemon]; !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	data_pokemon := cfg.Pokedex[pokemon]
	fmt.Printf("Name: %v\nHeight: %v\nStats:\n", data_pokemon.Name, data_pokemon.Height)

	for _, s := range data_pokemon.Stats {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range data_pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)
	}
	return nil
}

func CPoke(cfg *Config) error {
	fmt.Println("Your Pokedex:")
	for pokemon := range cfg.Pokedex {
		fmt.Printf("  - %s\n", pokemon)
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
