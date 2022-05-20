// To execute Go code, please declare a func main() in a package "main"

/*
Input: nums = [0,0,1,1,1,2,2,3,3,4]
Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

Example:
iterate thru array w one variable for the non-dup index
and the other searching forward for the first non-dup value

nonDupIdx:0
nextIdx:0
[0,0,1,1,1,2,2,3,3,4]

nonDupIdx:1
nextIdx:1
[0,0,1,1,1,2,2]
 0,1,
if val @ next == val @ nonDup then continue the loop
looking for the next nonDup

algorithm
write func removeDuplicates
- declare and init nonDupIdx to 0
- declare and init isDuplicate to false
loop over array starting at idx 1 stop when reach end
	- if currVal == nonDupVal && !isDuplicate
		- set isDuplicate to true
		- increment nonDupIdx
		- continue
	- if currVal != nonDupVal && isDuplicate
		- setNonDup to currVal
		- increase nonDupIdx by 1
- return nonDupIdx - 1

*/

package main

import "fmt"

/*
//First Attempt
func removeDuplicates(nums []int) int {
	nonDupIdx := 0
	isDuplicate := false

	if len(nums) < 2 {
		return 1
	}

	for idx := 1; idx < len(nums); idx++ {
		currVal := nums[idx]
		nonDupVal := nums[nonDupIdx]

		if isDuplicate {
			if currVal > nonDupVal {
				nums[nonDupIdx] = currVal
				nonDupIdx++
				nums[nonDupIdx] = currVal
			}
		} else {
			if currVal > nonDupVal {
				nonDupIdx++
			} else if currVal == nonDupVal {
				isDuplicate = true
				nonDupIdx++
			}
		}
	}

	if !isDuplicate {
		return nonDupIdx + 1
	}

	return nonDupIdx
}
*/
//Cleaned up Team Attempt
func removeDuplicates(nums []int) int {

	seen := nums[0]

	a := 0
	r := 0

	for r < len(nums) {
		if nums[r] == seen && r != 0 {
			r++
		} else {
			seen = nums[r]
			nums[a], nums[r] = nums[r], nums[a]
			r++
			a++
		}
	}

	return a
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

/*
a = 0
r = 1
if a == r {
	a++
	r++
} else if r > a {
	swap
	r++
}
*/
