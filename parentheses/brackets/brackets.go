// Package bracketsMap checks whether the passed string is a balanced sequence of bracketsMap.
package brackets

import (
	"errors"
	"math/rand"
	"time"
)

// Package level Error.
var ErrIncorrectLength = errors.New("incorrect Length, you need a positive number")

var (
	bracketsMap = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	bracketsArray = "(){}[]"
)

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
			if stackLength == 0 || r != bracketsMap[stack[stackLength-1]] {
				return false
			}
			stack = stack[:stackLength-1]
		}
	}

	return len(stack) == 0
}

// GenerateRandomSequence generates the random sequence of brackets with the specified Length.
func GenerateRandomSequence(length int) (string, error) {
	if length <= 0 {
		return "", ErrIncorrectLength
	}

	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	randomLimit := len(bracketsArray)

	for i := 0; i < length; i++ {
		bytes[i] = bracketsArray[rand.Intn(randomLimit)]
	}

	return string(bytes), nil
}
