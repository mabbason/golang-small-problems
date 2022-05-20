package main

/*
PROBLEM
Turn a path into a simplified path
    * input: string, representing path
    * output: string, representing simplified ^
    * rules:
        * path should start with '/'
        * double slashes should be removed
        * there should be no trailing '/'
        * remove '.' or '..'

EXAMPLES
* simplifyPath("/home/") --> "/home"
* simplifyPath(/../") --> "/"
* simplifyPath("/home//foo/") --> "/home/foo"

APPROACH
* Add values to a queue with rules (e.g. if adding a slash and the last item added to the queue is a slash, discard it)

ALGO
simplifyPath(path string) --> string
* declare `queue` and initialize to empty string slice
* declare `currentDir` and initialize to empty string
* declare `canonicalPath` and initialize to empty string
* iterate over characters in `path`
    * if current character is a slash, add `currentDir` to `queue` after checking rules (is not a '/', '.', or '..')
        * reset `currentDir` to empty string
    * add character to `currentDir`
* dequeue from `queue`, building up `canonicalPath` to return

*/

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(simplifyPath("/home//foo/")) // "/home/foo"
	fmt.Println(simplifyPath("/../"))        // "/""
}

func simplifyPath(path string) string {
	//split string
	dirs := strings.Split(path, "/")
	stack := &Stack{vals: []string{}}
	//iterate slice of path
	for _, dir := range dirs {
		if dir == "" || dir == "." {
			continue
		} else if dir == ".." {
			stack.pop()
		} else {
			stack.push(dir)
		}
	}
	/*
		 if currPathItem is NOT, "." or "" or ".."
			- push currPathItem to stack
		else if currPathItem is "" or "."
			- ignore/continue==
		else if currPathItem is ".."
			- pop top of stack off

		once path slice is DONE iterating
		declare and init canonicalPath to ""
	*/
	canonicalPath := ""

	for !stack.isEmpty() {
		path := *stack.pop()
		canonicalPath = "/" + path + canonicalPath
	}

	if len(canonicalPath) == 0 {
		return "/"
	}

	return canonicalPath

	/*
		for !stack.isEmpty() {
			val := stack.pop()
			canonicalPath += "\" + val
		}
		if len(canonicalPath) == 0 {
			return "\"
		}
		return canonicalPath
	*/

}

type Stack struct {
	vals []string
}

func (s *Stack) pop() *string {
	if !s.isEmpty() {
		val := s.vals[len(s.vals)-1]
		s.vals = s.vals[:len(s.vals)-1]
		return &val
	}
	return nil
}

func (s *Stack) push(str string) {
	s.vals = append(s.vals, str)
}

func (s *Stack) isEmpty() bool {
	return len(s.vals) == 0
}
