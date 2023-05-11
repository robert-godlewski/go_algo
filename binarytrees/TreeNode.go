package binarytrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (node *TreeNode) push(value int) {
	// Will not push if node.Val == value
	if node.Val > value {
		if node.Left == nil {
			node.Left = &TreeNode{Val: value, Left: nil, Right: nil}
		} else {
			node.Left.push(value)
		}
	} else if node.Val < value {
		if node.Right == nil {
			node.Right = &TreeNode{Val: value, Left: nil, Right: nil}
		} else {
			node.Right.push(value)
		}
	}
}
