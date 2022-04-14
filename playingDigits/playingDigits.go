package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(digPow(89, 1))    // 1
	fmt.Println(digPow(92, 1))    // -1
	fmt.Println(digPow(695, 2))   // 2
	fmt.Println(digPow(46288, 3)) // 2
}

/*
i: two ints, num and p
o: int result
r:
break the input digit into a "specialSum"
  - each digit in num to the progressive power of p... summed up
	- eg. given 695 as num and 2 as p
	- 6^(p) + 9^(p+1) + 5^(p+2) = specialSum
	- 6^2 + 9^3 + 5^4

Then use input number * successive ints k
to see if num * k = specialSum
if yes return k
if no (num *k > specialSum and no k was found) then return -1

algorithm
write func digPow takes two ints, num and p, returns int
- declare k and init to -1
- declare specialSum and init to 0
- declare stringNum and init as converted num to string
- iterate over stringNum
	- convert each strNum digit to numDigit
	- specialSum = specialSum + (numDigit^p)
	- increment p by 1
declare potentialK and init to 0
loop while num*potentialK is less than specialSum
	- if of num*potentialK is equal to specialSum
		- return k
	- else increment potentialK by 1
return k


*/

func digPow(num, p int) int {
	k := -1
	fP := float64(p)
	specialSum := 0.00
	stringNum := strconv.Itoa(num)
	for _, r := range stringNum {
		dig := float64(r - '0')
		specialSum += math.Pow(dig, fP)
		fP += 1
	}
	// println(int(specialSum))
	potentialK := 0
	for num*potentialK <= int(specialSum) {
		if num*potentialK == int(specialSum) {
			k = potentialK
		}
		potentialK += 1
	}

	return k
}
