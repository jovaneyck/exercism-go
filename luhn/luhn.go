package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func toDigits(number string) []int {
	result := make([]int, len(number))
	for i, d := range number {
		result[i], _ = strconv.Atoi(string(d))
	}
	return result
}

func doubleEvens(digits []int) {
	for i := len(digits) - 2; i >= 0; i = i - 2 {
		digits[i] = double(digits[i])
	}
}

func double(digit int) int {
	d := 2 * digit
	if d > 9 {
		return d - 9
	}
	return d
}

func sum(digits []int) int {
	result := 0
	for _, d := range digits {
		result += d
	}
	return result
}

func trim(s string) string {
	return strings.Replace(s, " ", "", -1)
}

func valid(input string) bool {
	for _, c := range input {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func allZeroes(digits []int) bool {
	for _, d := range digits {
		if d != 0 {
			return false
		}
	}
	return true
}

func Valid(number string) bool {
	trimmed := trim(number)
	if !valid(trimmed) {
		return false
	}
	digits := toDigits(trimmed)
	if allZeroes(digits) {
		return len(digits) > 1
	}

	doubleEvens(digits)
	s := sum(digits)
	return (s % 10) == 0
}
