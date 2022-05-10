package main

import (
	"fmt"
	"sort"
)

func main() {
	testIntervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}
	fmt.Println(merge(testIntervals))
}

func merge(intervals [][]int) [][]int {
	sort.SliceStable(intervals[:], func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		lastMerged := merged[len(merged)-1]
		if pairsOverlap(lastMerged, intervals[i]) {
			bothPairsSlice := []int{lastMerged[0], lastMerged[1],
				intervals[i][0], intervals[i][1]}
			min, max := findMinAndMax(bothPairsSlice)
			lastMerged[0], lastMerged[1] = min, max
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

func findMinAndMax(slice []int) (min int, max int) {
	min = slice[0]
	max = slice[0]
	for _, value := range slice {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func pairsOverlap(pair1, pair2 []int) bool {
	return pair1[1] >= pair2[0]
}
