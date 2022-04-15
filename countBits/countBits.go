package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(countBits(1234)) // prints 5
}



/*
Write a function that takes an integer as input, 
and returns the number of bits that are equal to 
one in the binary representation of that number. 
You can guarantee that input is non-negative.

Example: The binary representation of 1234 is 10011010010, 
so the function should return 5 in this case

i: int (positive)
o: int 0 - ?

algorithm
write func countBits takes int returns int
- convert input to binary number string
- declare and init count to 1
- iterate over string
	- if current char equals 1
		- increment count +1
RETURN count
*/

func countBits(num int64) int{
	binNum := strconv.FormatInt(num, 2)
	count := 0
	for _, r := range binNum {
		rToMatch := '1'
		if r == rToMatch {
			count += 1
		}
	}
	return count
}