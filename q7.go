package main

import (
	"euler/utils"
	"fmt"
)

func solveQ7(targetIndex int) int {
	num := 2
	primeIndex := 1

	for {
		if utils.IsPrime(num) == true {
			primeIndex++
		}

		if primeIndex > targetIndex {
			break
		}

		num++
	}

	return num
}

func main() {
	fmt.Println(solveQ7(6))
	fmt.Println(solveQ7(10001))
}
