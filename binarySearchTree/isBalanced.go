//isBalanced
func isBalanced(root *TreeNode) bool { //Leetcode func sig
	if root == nil {
		return true
	}
	return isBalanced(root.Left) && isBalanced(root.Right) && absOfDifference(maxDepth(root.Left), maxDepth(root.Right)) < 2
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func absOfDifference(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}