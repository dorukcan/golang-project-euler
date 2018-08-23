package main

import "fmt"

func solveQ1(target int) int {
	total := 0

	for num := 2; num < target; num++ {
		if num%3 == 0 && num%5 == 0 {
			total += num
		} else if num%3 == 0 {
			total += num
		} else if num%5 == 0 {
			total += num
		}
	}

	return total
}

func main() {
	fmt.Println(solveQ1(10))
	fmt.Println(solveQ1(1000))
}
