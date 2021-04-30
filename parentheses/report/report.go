// Package report calculates the percentage of balanced rows for specified length.
package report

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

// workerPoolProcess processes the required number of requests from the counter
// and writes the result to the results channel.
func workerPoolProcess(url string, counter <-chan int, results chan<- bool) {
	for range counter {
		result, err := makeRequest(url)
		if err != nil {
			log.Println(err.Error())
			results <- false
		}
		results <- brackets.IsBalanced(result)
	}
}

// Start starts calculating the percentage of balanced sequences.
func Start() {
	resultPool := make(chan bool, N)
	requestsNumber := N / WorkerPoolSize

	for length := 2; length <= 8; length += 2 {
		url := RequestURL + strconv.Itoa(length)
		balancedCounter := 0
		requestPool := make(chan int, requestsNumber)

		for i := 0; i < WorkerPoolSize; i++ {
			go workerPoolProcess(url, requestPool, resultPool)
		}

		for i := 0; i < N; i++ {
			requestPool <- i
		}
		close(requestPool)

		for i := 0; i < N; i++ {
			if <-resultPool {
				balancedCounter++
			}
		}

		fmt.Printf("length: %d\tpercentage of balanced sequences: %.1f %%\n", length, float64(balancedCounter)/N*100)
	}
}
