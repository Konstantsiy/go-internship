// Package brackets checks whether the passed string is a balanced sequence of brackets.
package brackets

// Map of brackets for comparison.
var brackets = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// IsBalanced verifies if the given string is a balanced sequence of brackets.
func IsBalanced(str string) bool {
	var stack []rune
	runes := []rune(str)

	for _, r := range runes {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			stackLength := len(stack)
			if stackLength == 0 || r != brackets[stack[stackLength-1]] {
				return false
			}
			stack = stack[:stackLength-1]
		}
	}

	return len(stack) == 0
}
