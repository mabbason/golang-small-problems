package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4})) //[[-1,-1,2],[-1,0,1]]
	fmt.Println(threeSum([]int{0, 0, 0, 0}))          // [[0,0,0]]
	fmt.Println(threeSum([]int{}))                    // []
	fmt.Println(threeSum([]int{0}))                   // []
	// fmt.Println(twoSum([]int{1,2,3,4}, 6))
}

func twoSum(nums []int, target int) [][]int {
	start := 0
	end := len(nums) - 1
	twos := [][]int{}

	for start < end {
		if (nums[start] + nums[end]) < target {
			start++
		} else if nums[start]+nums[end] > target {
			end--
		} else {
			twos = append(twos, []int{nums[start], nums[end]})
			start++
		}
		for start < end && start != 0 && nums[start] == nums[start-1] {
			start++
		}
	}
	return twos
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	triplets := [][]int{}

	var complements [][]int
	var triplet []int
	anchor := 0

	for anchor < len(nums)-2 {
		for anchor < len(nums)-2 && anchor != 0 && nums[anchor] == nums[anchor-1] {
			anchor++
		}
		complements = twoSum(nums[anchor+1:], -nums[anchor])
		for _, twos := range complements {
			triplet = append(twos, nums[anchor])
			triplets = append(triplets, triplet)
		}
		anchor++
	}
	return triplets
}
