### parentheses
Package parentheses implements a web service that allows you to generate a string with a 
length that is passed through the request URL parameter. The string consists of a sequence 
of brackets. The service uses the following functions: 

- IsBalanced - verifies if the given string is a balanced sequence of brackets;
- GenerateRandomSequence - generates the random sequence of brackets with the specified length.

The package allows you to calculate the percentage of balanced sequences
for the specified lengths (2, 4, 8) and print the results to stdout with the command:
```bigquery
go run report/report.go
```
