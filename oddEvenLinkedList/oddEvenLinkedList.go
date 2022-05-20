package main

import (
	"fmt"
)

func main() {
	ll := makeNLengthLinkedList(3)
	printList(ll)
	ll = oddEvenList(ll)
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

	currNode := head
	for currNode != nil {
		fmt.Println("node ", currNode.Val)
		currNode = currNode.Next
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

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd, even := head, head.Next
	evenHead := head.Next

	for even != nil && even.Next != nil {
		odd.Next = odd.Next.Next
		even.Next = even.Next.Next
		odd = odd.Next
		even = even.Next
	}
	odd.Next = evenHead

	return head
}
