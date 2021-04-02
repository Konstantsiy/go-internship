package strfuncs

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	testTable := []struct {
		Source   string
		Expected string
	}{
		{
			Source:   "programmatic test",
			Expected: "tset citammargorp",
		},
		{
			Source:   "12 13   15",
			Expected: "51   31 21",
		},
		{
			Source:   "君は背が高いです",
			Expected: "すでい高が背は君",
		},
		{
			Source:   "¿Está María en casa?",
			Expected: "?asac ne aíraM átsE¿",
		},
		{
			Source:   "Ես պիցցա եմ ուզում",
			Expected: "մւոզւո մե ացցիպ սԵ",
		},
		{
			Source:   "",
			Expected: "",
		},
	}

	for _, tc := range testTable {
		result := ReverseString(tc.Source)

		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %q, got %q", tc.Expected, result)
		}
	}
}

func ExampleReverseString() {
	fmt.Println(ReverseString("Hello, 世界"))
	// Output: 界世 ,olleH
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
			Source:   "ўстаньце, хлопцы!",
			Expected: "Ўстаньце, хлопцы!",
		},
		{
			Source:   "Ես պիցցա եմ ուզում",
			Expected: "Ես պիցցա եմ ուզում",
		},
		{
			Source:   "θέλω πίτσα",
			Expected: "Θέλω πίτσα",
		},
		{
			Source:   "私と一緒に散歩に出てくる",
			Expected: "私と一緒に散歩に出てくる",
		},
		{
			Source:   "",
			Expected: "",
		},
	}

	for _, tc := range testTable {
		result := CapitalizeString(tc.Source)

		if result != tc.Expected {
			t.Errorf("Incorrect result. Expect %q, got %q", tc.Expected, result)
		}
	}
}

func ExampleCapitalizeString() {
	fmt.Println(CapitalizeString("βγες μια βόλτα μαζί μου"))
	// Output: Βγες μια βόλτα μαζί μου
}
