// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	testMaxtrix := [][]int{  
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
 }
	fmt.Println(spiralOrder(testMaxtrix)) // [1,2,3,6,9,8,7,4,5]
}

import "fmt"

fmt.Println(spiralOrder[[1,2,3],[4,5,6],[7,8,9]])

/*
examples:
[0][0], 
[x][y] start by increment y
(x = 0) (y = 0),

[0][1],
(x)(y+1)

[0][2],
(x)(y+1) // until y = n, then increment x, n--

[1][2],
(x+1, y)

[2][2],
(x+1, y) // until x = m, then decrement y, m--

[2][1],
(x)(y-1) 

[2][0],
(x)(y-1) // until y = start, then decrement x, start++

[1][0],
(x-1)(y) // until x = y+1, then start all over again...

[1][1] // increment y until y = n
*/

/*
start indexes = [0][0]
OuterLoop {
  first loop run until y = (len of sub),

  second loop run until x = (len of input)),

  third loop run until y = 0,

  fourth loop run until x = 1,

  start indexes++
}
*/

func nextDirection(currentDirection string) string {
  if currentDirection == "TOP" {
    return "RIGHT"
  } else if currentDirection == "RIGHT" {
    return "BOTTOM"
  } else if currentDirection == "BOTTOM" {
    return "LEFT"
  } else {
    return "TOP"
  }
}

func spiralOrder(matrix [][]int) []int {
  values := make(int[], 0)
  direction := "TOP"

  for len(matrix) > 0 {
    values = append(values, removeTopOfSlice(matrix, direction)
    direction = nextDirection(direction) 
    
  }
}

func removeSlice(matrix [][]int, side string) []int {
  if side == "TOP" {
    line := removeTop(matrix)
    matrix = matrix[1:]
    return line
  } else if side == "RIGHT" {
    for _, line := range matrix {
      matrix[len(matrix[0])-1]
    }
  }
}

func removeTop(matrix [][]int) []int {

}



func spiralOrder(matrix [][]int) []int {
	// m := len(matrix)
	n := len(matrix[0])

	for stIdx := 0; stIdx < len(matrix) - 1; stIdx++{
		for y := stIdx; y < n; y++ {
			fmt.Println("y", y)
		}

	}
  return []int{1,2,3}
}
