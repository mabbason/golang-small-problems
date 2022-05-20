func findMinIdx(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if right-left <= 1 {
			if nums[left] > nums[right] {
				return right
			} else {
				return left
			}

		}
		if nums[left] < nums[right] {
			return left
		} else if nums[mid] < nums[left] {
			right = mid
		} else if nums[mid] > nums[left] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func search(nums []int, target int) int {
	minIdx := findMinIdx(nums)
	fmt.Println(minIdx)
	targetIdx := -1

	if nums[minIdx] == target {
		return minIdx
	} else if target < nums[minIdx] {
		return -1
	}

	if target > nums[len(nums)-1] {
		frontArr := nums[:minIdx]
		// fmt.Println("front", frontArr)
		targetIdx = binarySearchIdx(frontArr, target)
	} else {
		backArr := nums[minIdx:]
		// fmt.Println("back", backArr)
		targetIdx = binarySearchIdx(backArr, target)
		if targetIdx == -1 {
			return targetIdx
		} else {
			targetIdx += minIdx
		}

	}

	return targetIdx
}

func binarySearchIdx(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}