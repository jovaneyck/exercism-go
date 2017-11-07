package acronym

import (
	"strings"
	"unicode"
)

func isDelimiter(c rune) bool {
	return string(c) == "-" || unicode.IsSpace(c)
}

func Abbreviate(s string) string {
	words := strings.FieldsFunc(s, isDelimiter)

	result := ""
	for _, w := range words {
		firstLetter := w[0:1]
		result += strings.ToUpper(firstLetter)
	}
	return result
}
