package main

import (
	"fmt"
)

func testPalimdrome(num int) int {

	test := num
	reversed := 0
	var digit int

	for test > 0 {

		digit = test % 10
		reversed = reversed*10 + digit
		test = test / 10

	}
	if num == reversed {
		return 1
	} else {
		return 0
	}
}

func main() {

	var largestPal int
	var tempNum int
	var check int

	for i := 0; i <= 999; i++ {

		for j := 0; j <= 999; j++ {

			tempNum = i * j

			check = testPalimdrome(tempNum)

			if check == 1 {

				if tempNum > largestPal {
					largestPal = tempNum
				}

			}

		}

	}

	fmt.Printf("%d", largestPal)
}
