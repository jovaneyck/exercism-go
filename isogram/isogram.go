package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(s string) bool {
	encounteredLetters := make(map[rune]bool, len(s))
	for _, letter := range strings.ToLower(s) {
		if _, alreadyEncountered := encounteredLetters[letter]; unicode.IsLetter(letter) && alreadyEncountered {
			return false
		}

		encounteredLetters[letter] = true
	}
	return true
}
