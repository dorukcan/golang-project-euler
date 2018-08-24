/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names,
begin by sorting it into alphabetical order. Then working out the alphabetical value for each name,
multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53,
is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
 */

package main

import (
    "fmt"
    "io/ioutil"
    "sort"
    "strings"
)

func solveQ22(target int) int {
    var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    var names []string

    dat, err := ioutil.ReadFile("./files/p022_names.txt")
    if err != nil {
        panic(err)
    }

    content := strings.Replace(string(dat), "\"", "", -1)
    names = strings.Split(content, ",")
    sort.Strings(names)

    result := 0

    for index, name := range names {
        score := 0

        for position := range name {
            char := string(name[position])
            score += strings.Index(alphabet, char) + 1
        }

        score = score * (index + 1)
        result += score
    }

    return result
}

func main() {
    fmt.Println(solveQ22(1000))
}
