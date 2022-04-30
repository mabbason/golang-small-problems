package main

import "fmt"

func main() {
	test1 := []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3} //{pos: [3, 7], peaks: [6, 3]}
	// test2 := []int{1, 2, 2, 2, 1, 3, 2, 1, 3}          //{ pos: [1, 5], peaks: [2, 3] }
	test3 := []int{2, 14, 13, -5, -5, 6, 4, 9, 0, 1, 1}
	fmt.Println(PickPeaks(test1))
	// fmt.Println(PickPeaks(test2))
	fmt.Println(PickPeaks(test3))
}

/*
In this kata, you will write a function that returns the positions and the values of the "peaks"
(or local maxima) of a numeric array.

For example, the array arr = [0, 1, 2, 5, 1, 0] has a peak at position 3 with a value of 5
(since arr[3] equals 5).

The output will be returned as an object with two properties: pos and peaks. Both of these
properties should be arrays. If there is no peak in the given array, then the output should
be {pos: [], peaks: []}.

Example: pickPeaks([3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3])
should return {pos: [3, 7], peaks: [6, 3]} (or equivalent in other languages)

All input arrays will be valid integer arrays (although it could still be empty),
so you won't need to validate the input.

The first and last elements of the array will not be considered as peaks
(in the context of a mathematical function, we don't know what is after
and before and therefore, we don't know if it is a peak or not).

Also, beware of plateaus !!! [1, 2, 2, 2, 1] has a peak while [1, 2, 2, 2, 3]
and [1, 2, 2, 2, 2] do not. In case of a plateau-peak, please only return
the position and value of the beginning of the plateau.
For example: pickPeaks([1, 2, 2, 2, 1]) returns {pos: [1], peaks: [2]}
(or equivalent in other languages)
*/

/*
i: single slice/arr
o: map w format { pos: []int, peaks: []int }

r:
- peak: a higher value surrounded by lower values (could be plateau, eg 12221)
- position: the index of the first value "in" the peak

example

i: [1,2,2,2,1,3,2,1,3]
o: { pos: [1, 5], peaks: [1, 3] }

algorithm
write func takes single slice and returns map[string][]int
- declare and init variable result to output map format
- declare and init slice to [] // for potential peak
- iterate over input slice starting at index 1
	- if the value preceding current value < current value
		- replace potentialPeak's position and value w current index and value respectively
	- else if index > 1 && current value < potentialPeak value
		- write potentialPeak values to result map
return map

*/

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

func newPosPeaks() *PosPeaks {
	posPeaks := PosPeaks{}
	posPeaks.Pos, posPeaks.Peaks = []int{}, []int{}
	return &posPeaks
}

func PickPeaks(arr []int) PosPeaks {
	result := newPosPeaks()
	potentialPeak := make([]int, 0, 2)

	for i := 1; i < len(arr); i++ {
		prevVal, currVal := arr[i-1], arr[i]

		if prevVal < currVal {
			hasPotentialPeak := len(potentialPeak) == 2
			if hasPotentialPeak {
				potentialPeak[0], potentialPeak[1] = i, currVal
			} else {
				potentialPeak = append(potentialPeak, i, currVal)
			}

		} else if i > 1 && len(potentialPeak) == 2 && currVal < potentialPeak[1] {
			result.Pos = append(result.Pos, potentialPeak[0])
			result.Peaks = append(result.Peaks, potentialPeak[1])
			potentialPeak = make([]int, 0, 2)
		}
	}
	return *result
}
