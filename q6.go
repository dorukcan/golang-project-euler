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
