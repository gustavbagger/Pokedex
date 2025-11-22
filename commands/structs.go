package commands

import (
	"github.com/gustavbagger/Pokedex/internal"
)

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
	Config      *Config
}

type Config struct {
	Next     string
	Previous string
}

func Support(cfg *Config, cache *pokecache.Cache) map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func() error { return CExit(cfg) },
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func() error { return CHelp(cfg, cache) },
		},
		"map": {
			Name:        "map",
			Description: "Next 20 locations in Pokemon",
			Callback:    func() error { return CMap(cfg, cache) },
		},
		"mapb": {
			Name:        "mapb",
			Description: "Previous 20 locations in Pokemon",
			Callback:    func() error { return CMapb(cfg, cache) },
		},
	}
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
