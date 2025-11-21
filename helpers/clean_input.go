package helpers

import (
	"strings"
)

func CleanInput(text string) []string {
	lower := strings.ToLower(text)
	list := strings.Fields(lower)
	return list
}
