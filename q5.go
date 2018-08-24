/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
    "fmt"
    "math"

    "golang-project-euler/utils"
)

func solveQ5(limit int) int {
    numbers := utils.MakeRange(1, limit)

    result := 1
    divCount := 0

    for num := 2; num < limit; {
        found := false

        for index, target := range numbers {
            if target%num == 0 {
                found = true
                numbers[index] = int(target / num)
            }
        }

        if found == false {
            if divCount != 0 {
                result *= int(math.Pow(float64(num), float64(divCount)))
                divCount = 0
            }

            num++
        } else {
            divCount++
        }
    }

    return result
}

func main() {
    fmt.Println(solveQ5(10))
    fmt.Println(solveQ5(20))
}
