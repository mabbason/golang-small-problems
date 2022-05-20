/*Intersection of Two Arrays

Given two integer arrays nums1 and nums2, return an array of their intersection.
Each element in the result must be unique and you may return the result in any order.

Example 1:
Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2]

Example 2:
Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [9,4] or ([4,9])


Constraints:
1 <= nums1.length, nums2.length <= 1000
0 <= nums1[i], nums2[i] <= 1000


[2]

[1,2,2,1]
     P

[2,2]
   P

[9, 4]

[4,9,5]
   P

[9,4,8,4]
       P


[7, 8, 9, 9, 9, 10]
        				P

[9, 10]
     P


1,2
					p

7,8,9,10,
			p


algo:
- set pointer at 0 for both arrays
- Loop until a pointer hits the end of array
	- If pointer values !=, iterate second array
	- If equal, iterate both arrays
			Push value

algo2:
- iterate first array
 - push value to hash as 1
 - if value exists, continue Loop

- iterate second array
 - check hash for each value
 - if value exists, push into result array - decrement in hash
 - if value does not exist, continue
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(intersection([]int{1,2,2,1}, []int{2,2}))
	fmt.Println(intersection([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))
}

func intersection(arr1, arr2 []int) []int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	var result []int

	p1 := 0
	p2 := 0

	for p1 < len(arr1) && p2 < len(arr2) {
		fmt.Println(p1, p2)
		if arr1[p1] < arr2[p2] {
			p1++
		} else if arr2[p2] < arr1[p1] {
			p2++
		} else {
			result = append(result, arr1[p1])
			for p1 > 0 && p1 < len(arr1) && arr1[p1] == arr1[p1-1] {
				p1++
			}
			for p2 > 0 && p2 < len(arr2) && arr2[p2] == arr2[p2-1] {
				p2++
			}

		}
	}

	return result
}

// func intersection(arr1, arr2 []int) []int {
// 	var result []int
// 	hash := make(map[int]bool)

// 	for _, val := range arr1 {
// 		if _, ok := hash[val]; !ok {
// 			hash[val] = true
// 		} else {
// 			continue
// 		}
// 	}

// 	for _, val := range arr2 {
// 		if hash[val] {
// 			result = append(result, val)
// 			hash[val] = false
// 		} else {
// 			continue
// 		}
// 	}

// 	return result

// }
