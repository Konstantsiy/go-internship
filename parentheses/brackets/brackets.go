// Package bracketsMap checks whether the passed string is a balanced sequence of bracketsMap.
package brackets

// Map of bracketsMap for comparison.
var bracketsMap = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
}

// IsBalanced verifies if the given string is a balanced sequence of bracketsMap.
func IsBalanced(str string) bool {
	var stack []rune
	runes := []rune(str)

	for _, r := range runes {
		switch r {
		case '(', '[', '{':
			stack = append(stack, r)
		case ')', ']', '}':
			stackLength := len(stack)
			if stackLength == 0 || r != bracketsMap[stack[stackLength-1]] {
				return false
			}
			stack = stack[:stackLength-1]
		}
	}

	return len(stack) == 0
}
