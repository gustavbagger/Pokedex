package commands

import (
	"fmt"

	"github.com/gustavbagger/Pokedex/helpers"
	pokecache "github.com/gustavbagger/Pokedex/internal"
)

func Maps(url string, cfg *Config, cache *pokecache.Cache) error {

	data, err := helpers.RetrieveCache[LocationArea](url, cache)
	if err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Currently = url
	cfg.Previous = data.Previous

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}
