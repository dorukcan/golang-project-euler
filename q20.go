/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
 */

package main

import (
    "fmt"

    "golang-project-euler/utils"
)

func solveQ20(target int) int {
    fact := utils.FactorialNumberAsString(target)

    result := 0

    for index := range fact {
        result += utils.GetDigitFromString(fact, index)
    }

    return result
}

func main() {
    fmt.Println(solveQ20(10))
    fmt.Println(solveQ20(100))
}
