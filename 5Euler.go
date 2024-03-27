package main

import (
	"fmt"
)

func checkDivisable(num int) bool {

	for i := 1; i <= 20; i++ {

		if num%i != 0 {
			return false
		}
	}
	return true
}

func main() {
	for i := 2520; 1 == 1; i++ {

		if checkDivisable(i) == true {
			fmt.Println(i)
			return
		}
	}
}
