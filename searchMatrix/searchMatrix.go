// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

func main() {
	test1 := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	fmt.Println(searchMatrix(test1, 3))  // true
	fmt.Println(searchMatrix(test1, 13)) // false
}

/*
examples
step 1, binary search for the correct array


step 2, binary search for correct element
*/

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	last := len(matrix[0]) - 1
	targetArrIdx := 101

	// fmt.Println(left, right, last)

	for left <= right {
		mid := left + ((right - left) / 2)
		// fmt.Println("mid", mid)
		if matrix[mid][0] <= target && matrix[mid][last] >= target {
			// fmt.Println("matrixmid FIRST", matrix[mid][0], "matrixmid Last", matrix[mid][0], "target", target)
			targetArrIdx = mid
			break
		} else if target > matrix[mid][last] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if targetArrIdx == 101 {
		return false
	}
	arr := matrix[targetArrIdx]
	left, right = 0, len(arr)-1

	for left <= right {
		mid := left + ((right - left) / 2)
		if arr[mid] == target {
			return true
		} else if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	// fmt.Println("targArrIdx", targetArrIdx)

	return false
}
