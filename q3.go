package main

import (
	"euler/utils"
	"fmt"
	"math"
)

func main() {
	num := 600851475143

	maxDiv := int(math.Sqrt(float64(num))) + 1
	maxVal := 0

	for div := 2; div < maxDiv; div++ {
		if num%div == 0 && utils.IsPrime(div) == true {
			maxVal = div
		}
	}

	fmt.Println(maxVal)
}
