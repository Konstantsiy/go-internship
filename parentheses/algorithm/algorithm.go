// Package algorithm implements function for checking the balancing of parentheses in a string.
package algorithm

// Map of buckets for comparison.
var buckets = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// IsBalanced verifies if the given string is a balanced sequence of brackets.
func IsBalanced(str string) bool {
	var runes = []rune(str)
	var stack []rune

	for _, r := range runes {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			stackLength := len(stack)
			if stackLength == 0 || r != buckets[stack[stackLength-1]] {
				return false
			}
			stack = stack[:stackLength-1]
		}
	}

	return len(stack) == 0
}
