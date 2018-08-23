/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
	"fmt"
)

func solveQ9(target int) int {
	var a float64 = 1
	var b float64 = 1
	var c float64

	var target_ = float64(target)

	for a = 1; a < 1000; a++ {
		for b = 1; b < 1000; b++ {
			c = (target_ / 2) - (a * b / target_)
			if a*a+b*b == c*c && a+b+c == target_ {
				return int(a * b * c)
			}
		}
	}

	return 0
}

func main() {
	fmt.Println(solveQ9(1000))
}
