package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testTable := []struct {
		str      string
		expected string
	}{
		{
			str:      "Writing a programmatic test",
			expected: "tset citammargorp a gnitirW",
		},
		{
			str:      "12 13 114 15 16 -10",
			expected: "01- 61 51 411 31 21",
		},
		{
			str:      "localhost:5432",
			expected: "2345:tsohlacol",
		},
	}

	for _, testCase := range testTable {
		result := ReverseString(testCase.str)
		t.Logf("Calling ReverseString(%s), result: %s", testCase.str, result)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expected, result)
		}
	}
}

func TestCapitaliseString(t *testing.T) {
	testTable := []struct {
		source   string
		expected string
	}{
		{
			source:   "heLLO, wOrld!",
			expected: "Hello, World!",
		},
		{
			source:   "__Хабра__ хабр",
			expected: "__хабра__ Хабр",
		},
		{
			source:   "12345 some t ext",
			expected: "12345 Some T Ext",
		},
	}

	for _, testCase := range testTable {
		result := CapitalizeString(testCase.source)
		t.Logf("Calling CapitalizeString(%s), result: %s", testCase.source, result)
		if result != testCase.expected {
			t.Errorf("Incorrect result. Expect %s, got %s", testCase.expected, result)
		}
	}
}
