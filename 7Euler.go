package main

import (
	"fmt"
)

func findPrime(num int) bool {

	if num == 1 {

		return false

	} else if num <= 3 {

		return true

	} else if num%2 == 0 || num%3 == 0 {

		return false

	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {

			return false

		}
	}
	return true
}

func main() {
	//counter starts at 1 to catch 2 as a prime number
	counter := 0
	i := 1
	for {
		if findPrime(i) == true {
			counter++
			if counter == 10001 {
				fmt.Println(i)
				return
			}
		}
		i += 2
	}
}
