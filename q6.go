package main

import (
	"fmt"
	"math"
)

func main() {
	target := 100

	sumOfSquares := 0
	squareOfSum := 0

	for num := 1; num <= target; num++ {
		sumOfSquares += int(math.Pow(float64(num), 2))
		squareOfSum += num
	}

	squareOfSum = int(math.Pow(float64(squareOfSum), 2))

	fmt.Println(squareOfSum - sumOfSquares)
}
