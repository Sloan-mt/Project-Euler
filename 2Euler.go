package main

import (
	"fmt"
)

func main() {
	var value1 int = 1
	var value2 int = 2
	var value3 int
	var evenSum int = 2

	for value3 <= 4000000 {

		value3 := value2 + value1
		if value3%2 == 0 {
			evenSum += value3
		}

		value1 = value2
		value2 = value3

		if value3 >= 4000000 {
			break
		}
	}

	fmt.Printf("%d", evenSum)
}
