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
	// RequestURL is the url for service request.
	RequestURL = "http://localhost:8080/generate?n="
	// WorkerPoolSize is goroutines worker size.
	WorkerPoolSize = 50
)

// makeRequest processes a request with a specific URL.
func makeRequest(url string) (string, error) {
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

// Start starts calculating the percentage of balanced sequences.
func Start() {
	requestsNumber := N / WorkerPoolSize

	for length := 2; length <= 8; length += 2 {
		url := RequestURL + strconv.Itoa(length)
		balancedCounter := 0
		resultsPool := make(chan bool, WorkerPoolSize)

		for i := 0; i < WorkerPoolSize; i++ {
			go func() {
				for j := 0; j < requestsNumber; j++ {
					result, err := makeRequest(url)
					if err != nil {
						log.Println(err.Error())
						resultsPool <- false
					}
					resultsPool <- brackets.IsBalanced(result)
				}
			}()
		}

		for i := 0; i < N; i++ {
			if <-resultsPool {
				balancedCounter++
			}
		}

		fmt.Printf("length: %d\tpercentage of balanced sequences: %.1f %%\n", length, float64(balancedCounter)/N*100)
	}
}
