package main

import (
	"euler/utils"
	"fmt"
	"math"
)

func solveQ3(num int) int {
	maxDiv := int(math.Sqrt(float64(num))) + 1
	maxVal := 0

	for div := 2; div < maxDiv; div++ {
		if num%div == 0 && utils.IsPrime(div) == true {
			maxVal = div
		}
	}

	return maxVal
}

func main() {
	fmt.Println(solveQ3(13195))
	fmt.Println(solveQ3(600851475143))
}
