package helpers

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " Charmander Bulbasaur PIKACHU ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    " ",
			expected: []string{},
		},
		{
			input:    " Big SmALLs big",
			expected: []string{"big", "smalls", "big"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("wrong length")
			return
		}

		for i := range actual {
			word := actual[i]
			expected_word := c.expected[i]
			if word != expected_word {
				t.Errorf("Expected: %s, Got: %s", word, expected_word)
			}
		}

	}
}
