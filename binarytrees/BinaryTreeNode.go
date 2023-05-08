package binarytrees

type BinaryTreeNode struct {
	Val   int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func (node *BinaryTreeNode) Push(value int) {
	if node.Val > value {
		if node.Left == nil {
			node.Left = &BinaryTreeNode{Val: value, Left: nil, Right: nil}
		} else {
			node.Left.Push(value)
		}
	} else if node.Val < value {
		if node.Right == nil {
			node.Right = &BinaryTreeNode{Val: value, Left: nil, Right: nil}
		} else {
			node.Right.Push(value)
		}
	}
}
