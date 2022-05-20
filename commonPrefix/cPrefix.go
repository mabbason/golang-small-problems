package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"})) //"fl"
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))    // ""
}

/*
i: arr of strings
o: string, (longest common prefix)

r:
1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.

example
"flower","flow","flight" >> "fl" is common
"dog","racecar","car" >> "" (there is nothing that is common between any of the strings)

algorithm
write func longestCommonPrefix, takes a slice of strings, returns a string
- declare and init "longest" to empty string
- decalre and init "prefixes" to emtpy hash map, key type string, value type bool
- iterate over the input arr
	- iterate over the word string to grab each progressively longer prefix
		(eg. "dog" would be "d" >> "do" >> "dog" on each loop)
		- decalre and init "prefix" to the curr prefix
		- check if prefix exists in hashmap
			- YES
				- assign longest to curr prefix
			- NO
				- set value of hashmap key:prefix value to true
- return "longest"
*/

func longestCommonPrefix(strs []string) string {
	longest := ""
	prefixes := make(map[string]bool)

	if len(strs) == 1 {
		return strs[0]
	}

	for _, word := range strs {
		for i := 1; i < len(word)+1; i++ {
			prefix := word[:i]
			// fmt.Println(prefix, i)
			if prefixes[prefix] {
				longest = prefix
			} else {
				prefixes[prefix] = true
			}
		}
	}
	return longest
}

/*
!!!Need to work on algorithm:
Input:
["reflower","flow","flight"]
Output:
"fl"
Expected:
""
*/
