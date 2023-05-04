package queuesstacks

import (
	"fmt"
	"go_algo/linkedlists"
)

type LinkedStack struct {
	linkedlists.LinkedList
}

func (this *LinkedStack) Push(val int) {
	// Purpose is to add this only at the tail/top of the stack
	this.AddAtTail(val)
}

func (this *LinkedStack) Pop() int {
	val := this.Tail.Val
	this.Tail = this.Tail.Prev
	if this.Tail != nil {
		this.Tail.Next = nil
	}
	this.Size--
	return val
}

func (this *LinkedStack) PrintStack() string {
	result := "(bottom) -> "
	if this.Size == 0 {
		result += "none -> "
	} else {
		i := 0
		cur := this.Head
		for cur != this.Tail.Next && i < this.Size {
			result += fmt.Sprintf("%d", cur.Val)
			result += " -> "
			cur = cur.Next
			i++
		}
	}
	result += "(top)"
	return result
}
