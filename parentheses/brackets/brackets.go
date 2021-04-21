// Package bracketsMap checks whether the passed string is a balanced sequence of bracketsMap.
package brackets

import (
	"errors"
	"math/rand"
	"time"
)

// Package level Error.
var ErrIncorrectLength = errors.New("incorrect length when you need a positive number")

// Map of bracketsMap for comparison.
var (
	bracketsMap = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	bracketsArray = "()[]{}"
)

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

// GetRandomSequence generates the random sequence of brackets with the specified length.
func GetRandomSequence(length int) (string, error) {
	if length <= 0 {
		return "", ErrIncorrectLength
	}

	rand.Seed(time.Now().UnixNano())
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(bracketsArray[rand.Intn(length)])
	}

	return string(runes), nil
}
