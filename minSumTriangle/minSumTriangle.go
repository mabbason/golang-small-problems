// triangle
// 2          0
// 3 4        1
// 6 5 7      2
// 4 1 8 3    3

/*
    2
   5 6
 11 10 13
15 11 18 16

*/

// func minimumTotal(triangle [][]int) int {
//   mappedArr := [][]int{}

//   for row := 0; row < len(triangle); row++ {
//     if row == 0 {
//       mappedArr = append(mappedArr, []int{triangle[0][0]})
//       continue
//     }
//     mappedArr = append(mappedArr, []int{})
//     for col := 0; col < len(triangle[row]); col++ {
//       up := -1
//       upLeft := -1

//       if col == len(triangle[row])-1 {
//         up = mappedArr[row-1][col-1]
//       } else {
//         up = mappedArr[row-1][col]
//       }

//       if col == 0 {
//         upLeft = mappedArr[row-1][col]
//       } else {
//         upLeft = mappedArr[row-1][col-1]
//       }
//       currVal := triangle[row][col]
//       mappedVal := min(currVal+upLeft, currVal+up)
//       mappedArr[row] = append(mappedArr[row], mappedVal)

//     }
//   }
//   return min(mappedArr[len(mappedArr)-1]...)
// }

// func handleNegative(num int) int {
//   if num < 0 {}
// }

// func min(nums ...int)int {
// 	min := nums[0]

// 	for i:= 1; i < len(nums); i++ {
// 		if nums[i] < min {
// 			min = nums[i]
// 		}
// 	}
// 	return min
// }

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	DP := make(map[[2]int]int)
	lastRowOfDP := []int{}

	for row := 0; row < len(triangle); row++ {
		for col := 0; col < len(triangle[row]); col++ {
			// add value at current row and col in triagle + minimum of the value in DP that's at [row-1][col] and [row-1][col-1] to DP[row][col]
			// Exception: for the first value, insert it directly into DP at [0][0]
			if row == 0 && col == 0 {
				DP[[2]int{0, 0}] = triangle[row][col]
				continue
			}
			// check if the values directly above and above and to the left of the current row and col exist is DP. If they do not, assign that value to infinity
			val, ok := DP[[2]int{row - 1, col}]
			if !ok {
				val = math.MaxInt
			}

			val2, ok2 := DP[[2]int{row - 1, col - 1}]
			if !ok2 {
				val2 = math.MaxInt
			}

			DP[[2]int{row, col}] = triangle[row][col] + min(val, val2)
			if row == len(triangle)-1 {
				lastRowOfDP = append(lastRowOfDP, DP[[2]int{row, col}])
			}
		}
	}
	return minOfSlice(lastRowOfDP)
}

func minOfSlice(values []int) int {

	minVal := values[0]
	for _, val := range values {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}