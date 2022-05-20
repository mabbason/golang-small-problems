// El Brute
func twoSum(nums []int, target int) []int {
	indeces := make([]int, 2)
out:
	for outerIdx, outerVal := range nums {
		for innerIdx := outerIdx + 1; innerIdx < len(nums); innerIdx++ {
			if outerVal+nums[innerIdx] == target {
				indeces[0] = outerIdx
				indeces[1] = innerIdx
				break out
			}
		}
	}
	return indeces
}

// Optimized
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	output := make([]int, 2)

	for idx1, val1 := range nums {
		delta := target - val1
		idx2, present := m[delta]

		if present {
			output[0], output[1] = idx1, idx2
			break
		} else {
			m[val1] = idx1
		}
	}
	return output
}
