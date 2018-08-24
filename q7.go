/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
    "fmt"

    "golang-project-euler/utils"
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
