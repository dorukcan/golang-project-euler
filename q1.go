package main

import "fmt"

func main() {
	target := 1000
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

	fmt.Println(total)
}
