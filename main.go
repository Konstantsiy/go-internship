package main

import (
	"unicode"
)

// ReverseString to reverse a string.
func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// CapitalizeString to capitalizes a string.
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
