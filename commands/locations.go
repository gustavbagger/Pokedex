package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/gustavbagger/Pokedex/internal"
)

func Maps(url string, cfg *Config, cache *pokecache.Cache) error {
	var data LocationArea
	
	if entry,ok := cache.Get(url); ok {
		if err := json.Unmarshal(entry, &data); err != nil {
			return err
		}
	} else {
		rec, err := http.Get(url)
		if err != nil {
			return err
		}
		defer rec.Body.Close()

		info, err := io.ReadAll(rec.Body)
		if err != nil {
			return err
		}
		cache.Add(url,info)	

		if err := json.Unmarshal(info, &data); err != nil {
			return err
		}
	}
	
	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}
