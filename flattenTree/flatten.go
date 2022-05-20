/*
          1.   
         / \
       5    2
      /    /  \
     6    3    4
*/


/*
          1
         / \
       2    5
      / \    \
     3   4    6


		      1
           \
            2
           / \
          3   4
				//end of first loop, 2-3-4, 5-6 on stack
				
            2
           / \
          3   4

					  2
						 \
				      3
        //end of second loop
						
				    3
             \
						  4
			  //end of third loop
				
				    4
						 \
						  5
*/
STACK: 5 
N: 4

// ITERATIVE Method
func flatten(root *TreeNode) {
	if root == nil {
			return
	}
	
	stack := []*TreeNode{root}
	
	for len(stack) > 0 {
			// pop the stack
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			
			if n.Right != nil {
					stack = append(stack, n.Right)
			}
			
			if n.Left != nil {
					stack = append(stack, n.Left)   
			}
			
			if len(stack) > 0 {
					n.Right = stack[len(stack)-1] // peek the stack
			}
			
			n.Left = nil
	}
	
}

//Recursive Method
func flatten(root *TreeNode)  {
  if root == nil || (root.Left == nil && root.Right == nil) {
    return
  }

  left, right := root.Left, root.Right
  root.Left = nil

	if left != nil {
		root.Right = left
		rightLastNode := findLastRight(root.Right)
    rightLastNode.Right = right
	}
  
  flatten(root.Right)
}

func findLastRight(root *TreeNode) *TreeNode {
	if root.Right == nil {
			return root
	}
	return findLastRight(root.Right)
}