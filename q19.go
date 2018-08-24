/*
You are given the following information, but you may prefer to do some research for yourself.

1 Jan 1900 was a Monday.
Thirty days has September,
April, June and November.
All the rest have thirty-one,
Saving February alone,
Which has twenty-eight, rain or shine.
And on leap years, twenty-nine.
A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.

How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?
 */

package main

import "fmt"

func isSunday(day int) bool {
    if day % 7 == 0{
        return true
    }

    return false
}

func solveQ19() int {
    result := 0

    startYear := 1900

    for year := startYear+1; year <= 2000; year++ {
        leapDay := 0

        if year % 4 == 0{
            leapDay = 1
        }

        for month:=1; month <= 12 ; month++{
            monthSize := 31

            if month == 2{
                monthSize = 28 + leapDay
            } else if month == 4 || month == 6 || month == 9 || month == 11 {
                monthSize = 30
            }

            day := (year-startYear)*365 + month * monthSize + 1

            if isSunday(day){
                result++
            }
        }
    }

    return result
}

func main() {
    fmt.Println(solveQ19())
}
