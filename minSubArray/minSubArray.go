package main

import (
	"fmt"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

// func minSubArrayLen(target int, nums []int) int {
// 	minLen := 0

// 	for subLen := 1; subLen <= len(nums); subLen++ {
// 		for i := 0; i <= len(nums) - subLen; i++ {
// 			subArr := nums[i:i+subLen]
// 			if sum(subArr) >= target {
// 				if minLen == 0 {
// 					minLen = len(subArr)
// 				} else if len(subArr) < minLen {
// 					minLen = len(subArr)
// 				}
// 			}
// 		}
// 	}
// 	return minLen
// }

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func minSubArrayLen(target int, nums []int) int {
	minLen := 0
	a := 0

	for a < len(nums) {
		r := a + 1
		for r <= len(nums) {
			subArr := nums[a:r]
			if sum(subArr) >= target {
				if minLen == 0 {
					minLen = len(subArr)
				} else if len(subArr) < minLen {
					minLen = len(subArr)
				}
			}
			r++
		}
		a++
	}

	// for subLen := 1; subLen <= len(nums); subLen++ {
	//   for i := 0; i <= len(nums) - subLen; i++ {
	//     subArr := nums[i:i+subLen]
	//     if sum(subArr) >= target {
	//       if minLen == 0 {
	//         minLen = len(subArr)
	//       } else if len(subArr) < minLen {
	//         minLen = len(subArr)
	//       }
	//     }
	//   }
	// }
	return minLen
}

// target = 7
// bestYet = len(nums)+1
// [2,3,1,2,4,3]
//.           s (THEN) walk s foward until sum <= target, then walk e forward again
//.           e walk e forward until sum >=target
// assign bestYet the distance of e-s for length
//
//
//
//
//
//
