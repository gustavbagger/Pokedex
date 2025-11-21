package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Maps(url string, cfg *Config) error {
	rec, err := http.Get(url)
	if err != nil {
		return err
	}
	defer rec.Body.Close()

	var data LocationArea

	info, err := io.ReadAll(rec.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(info, &data); err != nil {
		return err
	}

	cfg.Next = data.Next
	cfg.Previous = data.Previous

	for _, result := range data.Results {
		fmt.Println(result.Name)
	}

	return nil
}
