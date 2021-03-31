package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	testTable := []struct {
		Source   string
		Expected string
	}{
		{
			Source:   "Writing a programmatic test",
			Expected: "tset citammargorp a gnitirW",
		},
		{
			Source:   "12 13 114 15 16 -10",
			Expected: "01- 61 51 411 31 21",
		},
		{
			Source:   "localhost:5432",
			Expected: "2345:tsohlacol",
		},
	}

	for _, tc := range testTable {
		result := ReverseString(tc.Source)

		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %s, got %s", tc.Expected, result)
		}
	}
}

func TestCapitalizeString(t *testing.T) {
	testTable := []struct {
		Source   string
		Expected string
	}{
		{
			Source:   "i am trying to break this code",
			Expected: "I am trying to break this code",
		},
		{
			Source:   "ўстаньце, хлопцы, ўстаньце, браткі! Ўстань ты, наша старана!",
			Expected: "Ўстаньце, хлопцы, ўстаньце, браткі! Ўстань ты, наша старана!",
		},
		{
			Source:   "",
			Expected: "",
		},
	}

	for _, tc := range testTable {
		result := CapitalizeString(tc.Source)

		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %s, got %s", tc.Expected, result)
		}
	}
}
