// To execute Go code, please declare a func main() in a package "main"

package main

import (
	"fmt"
)

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4})) // 5 or 1
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))          // 2
}

/*
	left, right = 0, last

  for left <= right {
    mid := left + ((right - left) / 2)
    if targetArr[mid] == target {
      return true
    } else if target > targetArr[mid] {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }
*/

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	if len(nums) == 1 {
		return 0
	} else if len(nums) == 2 {
		if nums[0] > nums[1] {
			return 0
		} else {
			return 1
		}
	}

	for left <= right {
		mid := left + ((right - left) / 2)
		fmt.Println()
		if mid == 0 {
			if nums[0] > nums[1] {
				return 0
			}
		} else if mid == len(nums)-1 {
			if nums[mid] > nums[mid-1] {
				return len(nums) - 1
			}
		}

		if nums[mid-1] < nums[mid] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid-1] < nums[mid] && nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}
