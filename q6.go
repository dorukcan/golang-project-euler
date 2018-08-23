/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import (
	"fmt"
	"math"
)

func solveQ6(target int) int {
	sumOfSquares := 0
	squareOfSum := 0

	for num := 1; num <= target; num++ {
		sumOfSquares += int(math.Pow(float64(num), 2))
		squareOfSum += num
	}

	squareOfSum = int(math.Pow(float64(squareOfSum), 2))

	return squareOfSum - sumOfSquares
}

func main() {
	fmt.Println(solveQ6(100))
}
