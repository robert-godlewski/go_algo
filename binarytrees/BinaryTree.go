package binarytrees

//import "fmt"

type BinaryTree struct {
	Root *BinaryTreeNode
	Size int
}

func Constructor() BinaryTree {
	return BinaryTree{Root: nil, Size: 0}
}

func (this *BinaryTree) Push(value int) {
	if this.Root == nil {
		this.Root = &BinaryTreeNode{Val: value, Left: nil, Right: nil}
	} else {
		this.Root.Push(value)
	}
	this.Size++
}
