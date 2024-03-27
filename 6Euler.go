package main

import (
	"fmt"
)

func getSumOfSquares() int {
	var sum int
	for i := 1; i <= 100; i++ {
		sum = i*i + sum
	}
	return sum
}

func getSquareOfSums() int {
	var sum int
	for i := 1; i <= 100; i++ {
		sum = i + sum
	}
	return sum * sum
}

func main() {
	sumOfSquares := getSumOfSquares()

	squareOfSums := getSquareOfSums()

	difference := squareOfSums - sumOfSquares

	fmt.Printf("%d", difference)
}
