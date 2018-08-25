/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number.
For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28,
which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less
than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16,
the smallest number that can be written as the sum of two abundant numbers is 24.
By mathematical analysis, it can be shown that all integers greater
than 28123 can be written as the sum of two abundant numbers.
However, this upper limit cannot be reduced any further by analysis even though
it is known that the greatest number that cannot be expressed
as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
 */

package main

import (
    "fmt"

    "golang-project-euler/utils"
)

func solveQ23() int {
    var abundants []int
    var sums []int

    upperLimit := 28123

    for num := 1; num<upperLimit; num++ {
        if utils.IsAbundantNumber(num) {
            abundants = append(abundants, num)

            for _, ab := range abundants {
                sums = append(sums, ab+num)
            }
        }
    }

    sums = utils.UniqueSlice(sums)

    result := 0

    for x:=1;x<upperLimit;x++ {
        found := false

        for _, s := range sums {
            if x == s {
                found = true
                break
            }
        }

        if !found {
            result += x
        }
    }

    return result
}

func main() {
    fmt.Println(solveQ23())
}
