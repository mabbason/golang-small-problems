package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	          1
	         / \
	       2    3
	      / \    \
	     4   5    6
		  /
		 7
*/

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, nil, nil}
	root.Right = &TreeNode{3, nil, nil}
	root.Left.Left = &TreeNode{4, nil, nil}
	root.Left.Right = &TreeNode{5, nil, nil}
	root.Right.Right = &TreeNode{6, nil, nil}
	root.Left.Left.Left = &TreeNode{7, nil, nil}

	fmt.Println("InOrder Traversal - recursive solution : ")
	InOrderTraverse(root)
	fmt.Println("")

	fmt.Println("PreOrder Traversal - recursive solution : ")
	PreOrderTraverse(root)
	fmt.Println("")

	fmt.Println("PostOrder Traversal - recursive solution : ")
	PostOrderTraverse(root)
	fmt.Println("")

	fmt.Println("LevelOrder Traversal - recursive solution : ")
	levelOrder(root)
	fmt.Println("")
}

func InOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraverse(root.Left)
	fmt.Println(root.Val)
	InOrderTraverse(root.Right)
}

func PreOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreOrderTraverse(root.Left)
	PreOrderTraverse(root.Right)
}

func PostOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraverse(root.Left)
	PostOrderTraverse(root.Right)
	fmt.Println(root.Val)
}

//LevelOrderTraverse
/*
Algorithm
1: Figure out the max depth of the tree
2: Loop for each level of the max depth, increment level each loop
	- on each loop, call function to print that current level, pass in the loop number

Function to Print Current Level
1: if root is null return
2: if level passed in is 1 that means it's the "current" level, so print value
3: else, call the function recursively with level-1 on each subtree
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	allLevels := [][]int{}
	treeHeight := maxDepth(root)
	for lvl := 1; lvl <= treeHeight; lvl++ {
		levelArr := []int{}
		levelArr = traverseCurrentLevel(root, lvl, levelArr)
		allLevels = append(allLevels, levelArr)
	}
	return allLevels
}

func traverseCurrentLevel(root *TreeNode, level int, levelArr []int) []int {
	if root == nil {
		return []int{}
	}
	if level == 1 {
		levelArr = append(levelArr, root.Val)
		fmt.Println(root.Val)
	}
	if level > 1 {
		traverseCurrentLevel(root.Left, level-1, levelArr)
		traverseCurrentLevel(root.Right, level-1, levelArr)
	}

	return levelArr
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}
