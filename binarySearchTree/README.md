       1
      / \
    2    3
   / \
  4   5
Depth First (Go all the way to the leaves, then return)

Inorder (Left, Root, Right)  - 4 2 5 1 3
  ```go
func traverseInOrder(root *TreeNode) {
  if root == nil {
    return
  }
	traverseInOrder(root.Left)
  fmt.Println(root.Val)
  traverseInOrder(root.Right)
}
  ```

Preorder (Root, Left, Right) - 1 2 4 5 3
  ```go
func traversePreOrder(root *TreeNode) {
  if root == nil {
    return
  }
	fmt.Println(root.Val)
  traverseInOrder(root.Left)
  traverseInOrder(root.Right)
}
  ```
Postorder (Left, Right, Root)- 4 5 2 3 1

Breadth-First
  Level Order (check both child nodes THEN move down a level)
    1 2 3 4 5

```go
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
func LevelOrderTraverse(root *TreeNode) {
	if root == nil {
		return
	}
	treeHeight := maxDepth(root)
	for lvl := 1; lvl <= treeHeight; lvl++{
		traverseCurrentLevel(root, lvl)
	} 
}

func traverseCurrentLevel(root *TreeNode, level int) {
	if (root == nil) {
		return
	}
	if (level == 1) {
		fmt.Println(root.Val)
	}	else if level > 1 {
		traverseCurrentLevel(root.Left, level - 1)
		traverseCurrentLevel(root.Right, level - 1)
	}
}

func max(a, b int) int{
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
```