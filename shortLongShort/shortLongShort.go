package main

import "fmt"

func main() {
	fmt.Println(shortLongShort("abc", "defgh")) // "abcdefghabc"
	fmt.Println(shortLongShort("", "xyz"))      // "xyz"
}

func shortLongShort(str1, str2 string) string {
	if len(str1) < len(str2) {
		return str1 + str2 + str1
	} else {
		return str2 + str1 + str2
	}
}
