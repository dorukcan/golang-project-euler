/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
    "fmt"

    "golang-project-euler/utils"
)

func solveQ10(target int) int {
    result := 0

    for num := 2; num < target; num++ {
        if utils.IsPrime(num) {
            result += num
        }
    }

    return result
}

func main() {
    fmt.Println(solveQ10(10))
    fmt.Println(solveQ10(2000000))
}
