// Package service implements a parentheses web service that
// generates a random sequence of parentheses of the given length.
package service

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Package level Error.
var ErrIncorrectLength = errors.New("incorrect length when you need a positive number")

const bracketsArray = "(){}[]"

// processRequest accepts a query with length parameter and
// returns the random generated string of brackets with specified length.
func processRequest(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")

	length, err := strconv.Atoi(n)
	if err != nil || length <= 0 {
		http.Error(w, ErrIncorrectLength.Error(), http.StatusBadRequest)
	}

	resultStr, err := generateRandomSequence(length)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Write([]byte(resultStr))
}

// generateRandomSequence generates the random sequence of brackets with the specified length.
func generateRandomSequence(length int) (string, error) {
	if length <= 0 {
		return "", ErrIncorrectLength
	}

	rand.Seed(time.Now().UnixNano())
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(bracketsArray[rand.Intn(length)])
	}

	return string(runes), nil
}

// Run starts starts handling requests.
func Run() {
	http.HandleFunc("/generate", processRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
