package main

import (
	"fmt"
	"math"
)

func main() {
	grid := [][]int{
		{1, 2, 5},
		{6, 4, 3},
		{8, 7, 9},
	}
	fmt.Println(minPathSum(grid)) // 19
}

/*
{1,2,5},
{6,4,3},
{8,7,9},

minPaths (as a grid)
{1, 3, 8}
{7, 7, 10},
{15, 14, 19}

*/

// func minPathSum(grid [][]int) int {
// 	mappedArr := [][]int{}

// 	for row := 0; row < len(grid); row++ {
// 			mappedArr = append(mappedArr, []int{})
// 			for col := 0; col < len(grid[0]); col++ {
// 					if row == 0 && col == 0 {
// 							mappedArr[0] = append(mappedArr[0], grid[0][0])
// 							continue
// 					}

// 					var up int
// 					var left int

// 					if row == 0 {
// 							up = math.MaxInt
// 					} else {
// 							up = mappedArr[row-1][col]
// 					}

// 					if col == 0 {
// 							left = math.MaxInt
// 					} else {
// 							left = mappedArr[row][col-1]
// 					}
// 					// fmt.Println("left:", left, "up:", up)
// 					// fmt.Print("min:", min(left, up))
// 					varMax := math.MaxInt + 1
// 					fmt.Println(varMax)

// 					current := grid[row][col]
// 					mappedVal := current + min(up, left)
// 					mappedArr[row] = append(mappedArr[row], mappedVal)

// 			}
// 	}
// 	return mappedArr[len(mappedArr)-1][len(mappedArr[0])-1]
// }

// func min(values ...int) int {
// minVal := values[0]
// for _, val := range values {
// 	if val < minVal {
// 		minVal = val
// 	}
// }
// return minVal
// }

func minPathSum(grid [][]int) int {
	minPaths := map[[2]int]int{}

	for rowNo, row := range grid {

		for colNo, val := range row {
			printGrid(hashToGrid(minPaths))
			if rowNo == 0 && colNo == 0 {
				minPaths[[2]int{rowNo, colNo}] = val
				continue
			}
			above, okAbove := minPaths[[2]int{rowNo - 1, colNo}]
			if !okAbove {
				above = math.MaxInt
			}
			left, okLeft := minPaths[[2]int{rowNo, colNo - 1}]
			if !okLeft {
				left = math.MaxInt
			}

			minPaths[[2]int{rowNo, colNo}] = val + min(above, left)
		}
	}

	printGrid(hashToGrid(minPaths))
	return minPaths[[2]int{len(grid) - 1, len(grid[0]) - 1}]
}

func hashToGrid(h map[[2]int]int) [3][3]int {
	grid := [3][3]int{}
	for k, v := range h {
		grid[k[0]][k[1]] = v
	}
	return grid
}

func printGrid(g [3][3]int) {
	for _, row := range g {
		fmt.Println(row)
	}
	fmt.Println("")
}

func min(values ...int) int {
	minVal := values[0]
	for _, val := range values {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
