/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
    "fmt"
    "math"

    "golang-project-euler/utils"
)

func solveQ4(digitCount int) int {
    start := math.Pow10(digitCount - 1)
    end := math.Pow10(digitCount) - 1

    maxVal := 0

    for first := start; first <= end; first++ {
        for second := start; second <= end; second++ {
            product := int(first * second)

            if utils.IsPalindrome(product) == true {
                if product > maxVal {
                    maxVal = product
                }
            }
        }
    }

    return maxVal
}

func main() {
    fmt.Println(solveQ4(2))
    fmt.Println(solveQ4(3))
}
