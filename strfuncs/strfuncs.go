// Package strfuncs provides functions for reverses and capitalizes strings.
package strfuncs

import (
	"unicode"
)

// ReverseString returns its argument string reversed rune-wise left to right.
func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// CapitalizeString converts the first character of an argument string to uppercase if it was in lowercase.
func CapitalizeString(str string) string {
	if len(str) == 0 {
		return str
	}
	runes := []rune(str)

	if unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}

	return string(runes)
}
