func search(nums []int, target int) int {

	if len(nums) >= 3 {
		for i, num := range nums {
			if num == target {
				return i
			}
		}
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[left] < nums[mid] && nums[left] <= target && target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}