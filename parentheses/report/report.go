// Package report calculates the percentage of balanced rows for specified length.
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Konstantsiy/go-internship/parentheses/brackets"
)

const (
	// N is a total number of requests.
	N = 1000
	//RequestURL is the url for service request.
	RequestURL = "http://localhost:8080/generate?n="
)

// MakeRequest processes a request with a specific URL.
func MakeRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("incorrect response status code")
	}

	return string(body), nil
}

// main starts calculating the percentage of balanced sequences.
func main() {
	for length := 2; length <= 8; length += 2 {
		url := RequestURL + strconv.Itoa(length)
		balancedCounter := 0

		for i := 0; i < N; i++ {
			sequence, err := MakeRequest(url)
			if err != nil {
				log.Println(err.Error())
			}
			if brackets.IsBalanced(sequence) {
				balancedCounter++
			}
		}

		fmt.Printf("length: %d\tpercentage of balanced sequences: %.1f %%\n", length, float64(balancedCounter)/N*100)
	}
}
