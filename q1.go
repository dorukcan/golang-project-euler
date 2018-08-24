/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

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
