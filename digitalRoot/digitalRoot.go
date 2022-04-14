package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(DigitalRoot(493193))
}

func DigitalRoot(n int) int {
	if n > 9 {
		strNum := strconv.Itoa(n)
		sumOfParts := 0
		for i := 0; i < len(strNum); i++ {
			digit := strNum[i : i+1]
			numPart, _ := strconv.Atoi(digit)
			sumOfParts += numPart
		}
		return DigitalRoot(sumOfParts)
	} else {
		return n
	}
}

/*
i: int
o: int, sum of all digits individually

example:
942 =  9+4+2 = 15 (still two digits) so... 1+5 = 6 , should return 6

algorithm
write func DigitalRoot takes single int as arg, returns single int
- if input num > 9
	- turn the num into a string
	- declare a sum and init to 0
	- iterate over the string
		- turn each letter char into a num
		- add the single char number to sum
		call func self (resursion) on the sum of the parts
- else return num
*/
