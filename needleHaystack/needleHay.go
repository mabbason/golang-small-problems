// To execute Go code, please declare a func main() in a package "main"

package main

import "fmt"

func main() {
	fmt.Println(strStr("hello", "ll"))                         // 2
	fmt.Println(strStr("hello world", " w"))                   // 5
	fmt.Println(strStr("hello to the martians on mars", "ar")) // 14
	fmt.Println(strStr("hello to the martians on mars", "rs")) // 27
	fmt.Println(strStr("hello", ""))                           // 0
	fmt.Println(strStr("hello", "helloWorld"))                 // -1
	fmt.Println(strStr("aaaaa", "bba"))                        // -1
}

func strStr(haystack string, needle string) int {
	for idx := 0; idx <= len(haystack)-len(needle); idx++ {
		if needle == haystack[idx:idx+len(needle)] {
			return idx
		}
	}
	return -1
}

/*
//Brute Force Method

func strStr(haystack string, needle string) int {
	for begIdx := 0; begIdx <= len(haystack) - len(needle); begIdx++ {
			for endIdx := begIdx + 1; endIdx <= begIdx + len(needle); endIdx++ {
					substr := haystack[begIdx:endIdx]

					if substr == needle {
							return begIdx
					}
			}
	}
	return -1
}
*/
