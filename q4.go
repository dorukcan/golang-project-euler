package main

import (
	"euler/utils"
	"fmt"
	"math"
)

func main() {
	digitCount := 3

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

	fmt.Println(maxVal)
}
