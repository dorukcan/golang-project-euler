package main

import (
	"euler/utils"
	"fmt"
)

func solveQ2(limit int) int {
	total := 0
	numbers := utils.Fib(limit)

	for _, num := range numbers {
		if num%2 == 0 {
			total += num
		}
	}

	return total
}

func main() {
	fmt.Println(solveQ2(100))
	fmt.Println(solveQ2(4000000))
}
