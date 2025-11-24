package helpers

import (
	"encoding/json"
	//"fmt"
	"io"
	"net/http"
	pokecache "github.com/gustavbagger/Pokedex/internal"
)


func RetrieveCache[T any](url string, cache *pokecache.Cache) (T, error) {

	var data T
	if entry, ok := cache.Get(url); ok {
		// fmt.Println("Retrieving from cache")
		if err := json.Unmarshal(entry, &data); err != nil {
			return data, err
		}
	} else {
		// fmt.Println("Retrieving from API")
		rec, err := http.Get(url)
		if err != nil {
			return data, err
		}
		defer rec.Body.Close()

		info, err := io.ReadAll(rec.Body)
		if err != nil {
			return data, err
		}
		cache.Add(url, info)

		if err := json.Unmarshal(info, &data); err != nil {
			return data, err
		}
	}
	return data,nil
}