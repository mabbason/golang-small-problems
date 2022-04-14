package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(spinWords("one"))                 // single word NOT reversed
	fmt.Println(spinWords("singleWord"))          // single word reversed
	fmt.Println(spinWords("two Words"))           //one not, one reversed
	fmt.Println(spinWords("hey feLLow Warriors")) //mixed case
}

/*
i: string "phrase" w one or more words
o: same string, but "processed"

r:
	- Processed looks like
		- if a word is of length > 5, reverse the word
		- else just keep the word as is

algorithm
write func spinWords takes single string as arg, returns a string
- split the input string into a slice
- iterate over the slice
	- for each value
		- if string len is greater than 4, reverse the string (helper -> reverseWord)
		- else continue to next loop
- join the slice back into a string
- return the result of the join
*/

func spinWords(str string) string {
	processed := []string{}
	words := strings.Split(str, " ")
	for _, word := range words {
		if len(word) > 4 {
			word = reverseWord(word)
		}
		processed = append(processed, word)
	}
	return strings.Join(processed, " ")
}

type Stack struct {
	chars []string
}

//push
func (s *Stack) Push(str string) {
	s.chars = append(s.chars, str)
}

//pop
func (s *Stack) Pop() string {
	len := len(s.chars)

	popped := ""
	if len != 0 {
		popped = s.chars[len-1]
	}

	if len > 1 {
		s.chars = s.chars[0 : len-1]
	} else {
		s.chars = []string{}
	}
	return popped
}

func reverseWord(str string) string {
	charsStack := Stack{}
	for _, c := range str {
		char := string(c)
		charsStack.Push(char)
	}

	reversedStr := ""
	for len(charsStack.chars) > 0 {
		popChar := charsStack.Pop()
		reversedStr += popChar
	}
	return reversedStr
}
