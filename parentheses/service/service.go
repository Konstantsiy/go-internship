// Package service implements a parentheses web service that
// generates a random sequence of parentheses of the given length.
package service

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Konstantsiy/go-internship/parentheses/brackets"
)

// ProcessRequest accepts a query with length parameter and
// returns the random generated string of brackets with specified length.
func ProcessRequest(w http.ResponseWriter, r *http.Request) {
	n := r.URL.Query().Get("n")

	length, err := strconv.Atoi(n)
	if err != nil || length <= 0 {
		http.Error(w, "Incorrect request param of string length. You need a positive number", http.StatusBadRequest)
		return
	}

	resultStr, err := brackets.GenerateRandomSequence(length)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write([]byte(resultStr))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Run starts handling requests.
func Run() {
	http.HandleFunc("/generate", ProcessRequest)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
