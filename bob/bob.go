package bob

import (
	"strings"
	"unicode"
)

func containsLetter(s string) bool {
	for _, letter := range s {
		if unicode.IsLetter(letter) {
			return true
		}
	}
	return false
}

func allUpper(s string) bool {
	return strings.ToUpper(s) == s
}

func Hey(remark string) string {
	trimmed := strings.TrimSpace(remark)
	switch {
	case trimmed == "":
		return "Fine. Be that way!"
	case containsLetter(remark) && allUpper(remark):
		return "Whoa, chill out!"
	case strings.HasSuffix(trimmed, "?"):
		return "Sure."
	default:
		return "Whatever."
	}
}
