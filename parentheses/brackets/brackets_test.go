package brackets

import (
	"regexp"
	"testing"
)

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("No error expected but got: %v", err)
	}
}

func assertError(t *testing.T, givenError, expectedError error) {
	t.Helper()
	if givenError != expectedError {
		t.Errorf("Eepected %v error but got another one: %v", expectedError, givenError)
	}
}

func TestIsBalanced(t *testing.T) {
	testTable := []struct {
		Str      string
		Expected bool
	}{
		{
			Str:      "({[]})",
			Expected: true,
		},
		{
			Str:      "(((",
			Expected: false,
		},
		{
			Str:      "{{{}()}} []([])",
			Expected: true,
		},
		{
			Str:      "(((1 + 2) * 3) - 4)/5",
			Expected: true,
		},
		{
			Str:      ")(",
			Expected: false,
		},
		{
			Str:      "345 + 434 * 4553",
			Expected: true,
		},
	}

	for _, tc := range testTable {
		result := IsBalanced(tc.Str)
		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %t, got %t", tc.Expected, result)
		}
	}
}

func TestGenerateRandomSequence(t *testing.T) {
	const regexPattern = "[^()[\\]{\\}]"

	testTable := []struct {
		Length          int
		IsErrorExpected bool
		ExpectedError   error
	}{
		{
			Length:          -12,
			IsErrorExpected: true,
			ExpectedError:   ErrIncorrectLength,
		},
		{
			Length:          0,
			IsErrorExpected: true,
			ExpectedError:   ErrIncorrectLength,
		},
		{
			Length:          4,
			IsErrorExpected: false,
			ExpectedError:   nil,
		},
		{
			Length:          100,
			IsErrorExpected: false,
			ExpectedError:   nil,
		},
	}

	for _, tc := range testTable {
		result, err := GenerateRandomSequence(tc.Length)

		if tc.IsErrorExpected {
			assertError(t, err, tc.ExpectedError)
		} else {
			assertNoError(t, err)
		}

		if matched, _ := regexp.Match(regexPattern, []byte(result)); matched {
			t.Errorf("The string must contain only parentheses")
		}
	}
}
