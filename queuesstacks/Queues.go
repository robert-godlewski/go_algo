package queuesstacks

import (
	"fmt"
	"go_algo/linkedlists"
)

type LinkedQueue struct {
	linkedlists.LinkedList
}

func (this *LinkedQueue) Push(val int) {
	// Puropse is to add this only at the tail of the queue
	this.AddAtTail(val)
}

func (this *LinkedQueue) Pop() int {
	val := this.Head.Val
	this.Head = this.Head.Next
	if this.Head != nil {
		this.Head.Prev = nil
	}
	this.Size--
	return val
}

func (this *LinkedQueue) PrintQueue() string {
	result := "(front) -> "
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
	result += "(back)"
	return result
}
