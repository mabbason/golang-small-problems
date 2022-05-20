package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValid("(){}[]")) // true
	fmt.Println(isValid("({[]})")) // true
	fmt.Println(isValid("()}[]"))  // false
	fmt.Println(isValid(""))       // true ???
}

var matchingParens map[string]string = map[string]string{
	"(": ")",
	"{": "}",
	"[": "]",
}

func isValid(s string) bool {
	stack := []string{}
	for i := 0; i < len(s); i++ {
		currentChar := s[i : i+1]
		if _, ok := matchingParens[currentChar]; ok {
			stack = append(stack, currentChar)
		} else if len(stack) == 0 || currentChar != matchingParens[stack[len(stack)-1]] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	//declare and init stack as empty string

	//iterate through input string

	// if currChar is opening bracket, push to stack
	//else if currChar is CLOSING bracket, and != top of stack OR stack is empty, return false
	//else pop from stack

	//return stack.isEmpty
	return len(stack) == 0
}
