/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import (
    "fmt"

    "golang-project-euler/utils"
)

func solveQ16(target int) int {
    digits := utils.PowStringsAsNumbers("2", target)
    result := 0

    for index := range digits{
        result += utils.GetDigitFromString(digits, index)
    }

    return result
}

func main() {
    fmt.Println(solveQ16(15))
    fmt.Println(solveQ16(1000))
}
