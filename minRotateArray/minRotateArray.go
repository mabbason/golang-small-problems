func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if right-left <= 1 {
			if nums[left] > nums[right] {
				return nums[right]
			} else {
				return nums[left]
			}

		}
		if nums[left] < nums[right] {
			return nums[left]
		} else if nums[mid] < nums[left] {
			right = mid
		} else if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			return nums[mid]
		}
	}
	return -1
}