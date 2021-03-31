package main

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
	runes := []rune(str)

	if (runes[0] >= 'a' && runes[0] <= 'z') || (runes[0] >= 'а' && runes[0] <= 'я') {
		runes[0] -= 32
	}

	return string(runes)
}
