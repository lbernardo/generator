package strings

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func ToPascalCase(input string) string {
	c := cases.Title(language.English)
	if strings.Contains(input, "_") {
		words := strings.Split(input, "_")
		for i, word := range words {
			words[i] = c.String(word)
		}
		return strings.Join(words, "")
	}

	words := strings.Fields(input)
	for i, word := range words {
		words[i] = c.String(word)
	}
	return strings.Join(words, "")
}
