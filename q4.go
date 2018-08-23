package main

import (
	"euler/utils"
	"fmt"
	"math"
)

func solveQ4(digitCount int) int {
	start := math.Pow10(digitCount - 1)
	end := math.Pow10(digitCount) - 1

	maxVal := 0

	for first := start; first <= end; first++ {
		for second := start; second <= end; second++ {
			product := int(first * second)

			if utils.IsPalindrome(product) == true {
				if product > maxVal {
					maxVal = product
				}
			}
		}
	}

	return maxVal
}

func main() {
	fmt.Println(solveQ4(2))
	fmt.Println(solveQ4(3))
}
