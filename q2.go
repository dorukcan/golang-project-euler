package main

import (
	"euler/utils"
	"fmt"
)

func main() {
	total := 0
	numbers := utils.Fib(4000000)

	for _, num := range numbers {
		if num%2 == 0 {
			total += num
		}
	}

	fmt.Println("total", total)
	fmt.Println("numbers", numbers)
}
