
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	curr := head
	next := curr.Next
	var prev *ListNode
	// var next *ListNode
	var temp *ListNode
	var subTail *ListNode
	nodeCount := 1

	for nodeCount < left {
		nodeCount++
		prev = curr
		curr = next
		next = next.Next
		fmt.Println("walk")
	}
	fmt.Println("nodeCount", nodeCount)
	fmt.Println("curr", curr.Val, "next", next.Val)
	temp = prev // last node before left
	subTail = curr
	// nodeCount = walk(nodeCount, prev, curr, next)

	//walk manually
	nodeCount++
	prev = curr
	curr = next
	next = next.Next
	fmt.Println("walk")

	for nodeCount <= right { //start when previous is on L and stop when previous is on R
		fmt.Println("nodeCount", nodeCount)
		// fmt.Println("prev", prev.Val, "curr", curr.Val, "next", next.Val)
		curr.Next = prev //rewire

		//walk manually
		nodeCount++
		prev = curr
		curr = next
		if next != nil {
			next = next.Next
		}
		fmt.Println("walk")
	}

	// //point temp to last reversed
	temp.Next = prev
	subTail.Next = curr

	return head
}



package main

import (
  "fmt"
)

func main() {
  ll := makeLinkedList(1,4)
  printList(ll)

  reverseBetween(ll, 2,4)

  printList(ll)
}


 type ListNode struct {
  Val int
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


func makeLinkedList(min, max int) *ListNode {
  currValue := min
  firstNode := &ListNode{
    Val: min,
  }
  var currNode = firstNode
  for currValue <= max {
    currValue++
    currNode.Next = &ListNode{
      Val: currValue,
    }
    currNode = currNode.Next
  }
  return firstNode
}

// func reverseBetween(head *ListNode, left int, right int) *ListNode {
//   dummy := &ListNode{}
//   dummy.Next = head
// 	prev := dummy

//   for i := 0; i < left - 1; i++ { 
//     prev = prev.Next 
//   }

// 	curr := prev.Next

//   for i := 0; i < right - left; i++ {
//       printList(dummy)
//       next := curr.Next
//       // fmt.Println("prev", prev.Val, "cur", cur.Val, "next", next.Val)
//       curr.Next = next.Next // like curr.Next.Next
//       next.Next = prev.Next
// 			prev.Next = next
//   }
  
//   return dummy.Next
// }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
  dummy := &ListNode{}
  dummy.Next = head
  preLeft := dummy
  for i := 0; i < left - 1; i++ { 
    preLeft = preLeft.Next // preLeft is moved into position before left
  }

  var prev *ListNode
	cur := head.Next
  
  for n := 0; n <= right - left; n++ { //iterate n times how many nodes to reverse
      next := cur.Next
      // fmt.Println("prev", prev.Val, "cur", cur.Val, "next", next.Val)
      cur.Next = prev
      prev = cur
      cur = next
  }
  head.Next.Next = cur
  head.Next = prev
  return dummy.Next
}