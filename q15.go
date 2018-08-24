/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?
 */

package main

import "fmt"

func calculateGrid(grid [][]int, gridLength int, x int, y int) int{
    if x == gridLength || y == gridLength{
        return 1
    } else {
        if grid[x][y] == 0 {
            right := calculateGrid(grid, gridLength, x+1, y)
            bottom := calculateGrid(grid, gridLength, x, y+1)

            grid[x][y] = right+bottom
        }

        return grid[x][y]
    }
}

func solveQ15(target int) int {
    grid := make([][]int, target)

    for i := range grid {
        grid[i] = make([]int, target)
    }

    return calculateGrid(grid, target, 0, 0)
}

func main() {
    fmt.Println(solveQ15(2))
    fmt.Println(solveQ15(20))
}
