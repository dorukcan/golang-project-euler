package main

import (
	"euler/utils"
	"fmt"
	"math"
)

func solve(limit int) int {
	numbers := utils.MakeRange(1, limit)

	result := 1
	divCount := 0

	for num := 2; num < limit; {
		found := false

		for index, target := range numbers {
			if target%num == 0 {
				found = true
				numbers[index] = int(target / num)
			}
		}

		if found == false {
			if divCount != 0 {
				result *= int(math.Pow(float64(num), float64(divCount)))
				divCount = 0
			}

			num++
		} else {
			divCount++
		}
	}

	return result
}

func main() {
	fmt.Println(solve(10))
	fmt.Println(solve(20))
}
