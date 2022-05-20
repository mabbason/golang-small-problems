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
*/

/*
Solution 1: Brute Force - Nested Loops; time(n^2)
   main function
    nested loops
      if the val in one arr == curr val in the other arr
        if isUnique(resultArr, newVal)
            add to resultArr
     return resultArr

   helper function isUnique, takes arr, and value
     set result to true
     iterate over the arr
        if arr at curr == val
            early return false
     return true

Solution 2: Hash Map; time(n+m)
    convert the larger array to hash (or both to hash)
    iterate over the smaller array and check the larger if present
        if NOT, append to result arr
    return result arr

Solution 3: Sort, then two Pointers; time(nlogn)
    sort both arrs
		pointer starting at the beginning of both arrays ie p1, p2
		if p1 == p2, add to result then
			increment both to new unique values

		Compare to the Previous Value within that Array itself
		Move the Pointer that is referering to a Smaller Value if p1 !== p2

		< if p2 Value is greater than the last value in P1 >

		Get the Last Index - Stopping Point
			- increment both pointers until they are greater/equal to stop point

N Log N + M Log M

nums1 = 1, 4,   5,  8, 9, 10, 11, 16
				        									p1
nums2 = 3, 4,   4,   8,   9,  9,  13, 14, 15, 17, 19, 21
				            													p2

				result [4, 8, 9]
*/