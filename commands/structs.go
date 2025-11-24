package commands

import pokecache "github.com/gustavbagger/Pokedex/internal"

type CliCommand struct {
	Name        string
	Description string
	Callback    func(string) error
	Config      *Config
}

type Config struct {
	Next      string
	Currently string
	Previous  string
	Pokedex   map[string]Pokemon
}

func Support(cfg *Config, cache *pokecache.Cache) map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    func(string) error { return CExit(cfg) },
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    func(string) error { return CHelp(cfg, cache) },
		},
		"map": {
			Name:        "map",
			Description: "Next 20 locations in Pokemon",
			Callback:    func(string) error { return CMap(cfg, cache) },
		},
		"mapb": {
			Name:        "mapb",
			Description: "Previous 20 locations in Pokemon",
			Callback:    func(string) error { return CMapb(cfg, cache) },
		},
		"explore": {
			Name:        "explore",
			Description: "List Pokemon at <location>",
			Callback:    func(place string) error { return CExpl(cfg, cache, place) },
		},
		"catch": {
			Name:        "catch",
			Description: "Try catching <pokemon>",
			Callback:    func(pokemon string) error { return CCat(cfg, cache, pokemon) },
		},
		"inspect": {
			Name:        "inspect",
			Description: "Shows details of <pokemon> in Pokedex",
			Callback:    func(pokemon string) error { return CInsp(cfg, cache, pokemon) },
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "List of pokemon in Pokedex",
			Callback:    func(string) error { return CPoke(cfg) },
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

type Place struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon        Pokemon `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:stat`
	} `json:stats`
	Types []struct {
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}
