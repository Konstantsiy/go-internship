package brackets

import (
	"testing"
)

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
