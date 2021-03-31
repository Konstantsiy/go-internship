package main

import (
	"fmt"
)

func ReverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func CapitalizeString(str string) string {
	runes := []rune(str)
	startWordFlag := true
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			startWordFlag = true
		} else {
			if ((runes[i] >= 'a' && runes[i] <= 'z') || (runes[i] >= 'а' && runes[i] <= 'я')) && startWordFlag {
				runes[i] -= 32
				startWordFlag = false
			} else if ((runes[i] >= 'A' && runes[i] <= 'Z') || (runes[i] >= 'А' && runes[i] <= 'Я')) && !startWordFlag {
				runes[i] += 32
			} else {
				startWordFlag = false
			}
		}
	}
	return string(runes)
}

func main() {
	s := "_HELLO_ WorLd"
	fmt.Println(ReverseString(s))
	fmt.Println(CapitalizeString(s))
}
