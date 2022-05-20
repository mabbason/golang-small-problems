package main

import (
	"fmt"
)

func main() {
	ll := makeNLengthLinkedList(4)
	printList(ll)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("Empty!")
	}
	nodeNumber := 1
	currNode := head
	for currNode != nil {
		fmt.Println("node #", nodeNumber, " value: ", currNode.Val)
		currNode = currNode.Next
		nodeNumber++
	}
	fmt.Println("")
}

func makeNLengthLinkedList(n int) *ListNode {
	if n == 0 {
		return nil
	}

	firstNode := &ListNode{
		Val: 1,
	}
	var currNode = firstNode
	for i := 2; i <= n; i++ {
		currNode.Next = &ListNode{
			Val: i,
		}
		currNode = currNode.Next
	}
	return firstNode
}
