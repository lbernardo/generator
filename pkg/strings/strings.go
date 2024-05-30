package strings

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func ToPascalCase(input string) string {
	c := cases.Title(language.English)
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = c.String(word)
	}
	return strings.Join(words, "")
}
