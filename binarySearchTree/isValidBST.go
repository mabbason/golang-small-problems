func isValidBST(root *TreeNode) bool {
	min := -math.MaxInt64
	max := math.MaxInt64
	return isValidBSTHelper(root, min, max)
}

func isValidBSTHelper(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	return root.Val > min && root.Val < max && isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}

/*
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Subtree problem statment

*/