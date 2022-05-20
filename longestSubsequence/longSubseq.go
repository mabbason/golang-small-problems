func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	if text1[0] == text2[0] {
		return 1 + longestCommonSubsequence(text1[1:], text2[1:])
	} else {
		return max(
			longestCommonSubsequence(text1, text2[1:]),
			longestCommonSubsequence(text1[1:], text2))
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func helper () {

// }

/*


abcd
ca


abcde
  |
"ace"
"abcde"

1) choose "ace" and hash characters to index position
string1hash = {

	"d": [0]
	"a": 1,
	"c": 2,
	"e": 3,
	"z": [4,5,6]
}

g t a b c d e a z z z

{
	0: {length: 0, indexFromLastString: nil}
	1: {length: 0, indexFromLast}
}


c
{
	c: [0,3]
	d: [1]
	a: [2]
	e: [4]
}

edcba

maxSeenSoFar 1
"c" in string 2
{
	0: {length: 1, indexSeenFromPreviousString: 2}
  1: {length: 1, indexSeenFromPreviousString: 1}
	2: {length: 2, indexSeenFromPreviousString: 3}
}

2) iterate over remaining string ("abcde") and at each character, determine if it is part of a common subsequence with the first string using the previously created hash. If it is not, carry forward the longest subsequence value from the previous index. If it is, add to the previous longest subsequence value

"abcde"
2)a. start with "a", which is at index 0. Check if "a" is in string1hash. It is, so we know we have a subsequence of length 1 for the substring of "abcde" "a", which occurs at index 0
newHash = {
	0: 1

}

2b. move to "b", which is at index 1. Check if "b" is in string1hash. It is not, so our new string "ab" still has a subsequence of length 1

newHash = {
	0: 1
	1: 1
}

2c. move to "c", which is at index 2. Check if "c" is in string1hash. It is, so we know we have a subsequence of length newHash[idx-1] + 1. Add this value to newHash[idx]
newHash = {
	0: 1,
	1: 1,
	2: 2

	..
}

...

retur newHash[len(string2)]

*/

/*
i: two strings
o: int of LENGTH of longest common subsequence

r -
finding if there is a common subsequence

*/

/*
  A   B   C   D
A
C
E
*/

// func longestCommonSubsequence(text1 string, text2 string) int {
//   dp := make(map[[2]int]int)

//   dp[[2]int{0,0}] = 1

//   for r := 0; r <= row; r++ {
//     for c := 0; c <= col; c++ {
//       if r == 0 && c == 0 {
//         continue
//       }

//       if obstacleGrid[r][c] == 1 {
//         dp[[2] int{r, c}] = 0
//         continue
//       }

//       dp[[2] int{r, c}] = dp[[2]int{r - 1, c}] + dp[[2]int{r, c - 1}]
//     }
//   }
//   return dp[[2]int {row, col}]
// }