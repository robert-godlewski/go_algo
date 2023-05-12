package binarytrees

import "fmt"

type BinaryTree struct {
	Root *TreeNode
	Size int
}

func Constructor() BinaryTree {
	return BinaryTree{Root: nil, Size: 0}
}

func (this *BinaryTree) Push(value int) {
	if this.Root == nil {
		this.Root = &TreeNode{Val: value, Left: nil, Right: nil}
	} else {
		this.Root.push(value)
	}
	this.Size++
}

func (this *BinaryTree) BST(value int) bool {
	result := false
	if this.Root != nil {
		cur := this.Root
		for cur != nil && result == false {
			fmt.Printf("cur.Val = %v\n", cur.Val)
			if cur.Val > value {
				cur = cur.Left
			} else if cur.Val < value {
				cur = cur.Right
			} else {
				// cur.Val == value
				result = true
			}
		}
	}
	return result
}
