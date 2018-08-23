package main

import (
	"euler/utils"
	"fmt"
)

func main() {
	targetIndex := 10001

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

	fmt.Println(num)
}
