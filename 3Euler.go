package main

import (
	"fmt"
)

func main() {
	var value int = 600851475143
	var largestPrime int
	//var num int

	for i := 2; i <= value; i++ {

		if value%i == 0 {
			value /= i
			if i > largestPrime {
				largestPrime = i
			}

		}

		if i > value {
			break
		}

	}
	fmt.Printf("%d\n", largestPrime)

}
