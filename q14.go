/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import "fmt"

func solveQ14(target int) int {
    maxChain := 0
    maxNum := 0

    for num:=2;num<target;num++{
        chainLength := 0

        for temp := num; temp != 1;{
            if temp % 2 == 0{
                temp = temp / 2
            } else {
                temp = 3* temp + 1
            }

            chainLength++
        }

        if chainLength > maxChain{
            maxChain = chainLength
            maxNum = num
        }
    }

    return maxNum
}

func main() {
    fmt.Println(solveQ14(1000000))
}
