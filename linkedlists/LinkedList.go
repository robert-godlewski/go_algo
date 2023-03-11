package linkedlists

import "fmt"

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
	Size int
}

func Constructor() LinkedList {
	return LinkedList{Head: nil, Tail: nil, Size: 0}
}

func (this *LinkedList) Get(index int) int {
	if this.Size == 0 {
		return -1
	}
	var cur ListNode = *this.Head
	for i := 0; i < this.Size; i++ {
		if i == index {
			return cur.Val
		}
		cur = *cur.Next
	}
	return -1
}

func (this *LinkedList) AddAtHead(val int) {
	var n = new(ListNode)
	n.Val = val
	if this.Head != nil {
		n.Next = this.Head
		this.Head.Prev = n
	}
	if this.Tail == nil {
		this.Tail = n
	}
	this.Head = n
	this.Size++
}

func (this *LinkedList) AddAtTail(val int) {
	var n = new(ListNode)
	n.Val = val
	if this.Tail != nil {
		n.Prev = this.Tail
		this.Tail.Next = n
	}
	if this.Head == nil {
		this.Head = n
	}
	this.Tail = n
	this.Size++
}

func (this *LinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.Size-1 {
		this.AddAtTail(val)
	} else if index < this.Size-1 {
		var n = new(ListNode)
		n.Val = val
		var prev = *this.Head
		var cur ListNode = *this.Head
		i := 0
		for i < index {
			if prev != cur {
				prev = cur
			}
			cur = *cur.Next
		}
		n.Prev = &prev
		n.Next = &cur
		prev.Next = n
		cur.Prev = n
		this.Size++
	}
}

func (this *LinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.Head = this.Head.Next
		this.Head.Prev = nil
	} else if index == this.Size-1 {
		this.Tail = this.Tail.Prev
		this.Tail.Next = nil
	} else if index < this.Size-1 {
		var prev = *this.Head
		var cur = *this.Head
		i := 0
		for i < index {
			if prev != cur {
				prev = cur
			}
			cur = *cur.Next
		}
		prev.Next = &cur
		cur.Prev = &prev
	}
	if index <= this.Size-1 {
		this.Size--
	}
}

func (this *LinkedList) PrintLL() string {
	result := "(head) -> "
	if this.Size == 0 {
		result += "none"
	} else {
		// Need to fix this section, returns an error for some reason
		i := 0
		var cur ListNode = *this.Head
		for cur != *this.Tail.Next && i < this.Size {
			result += fmt.Sprintf("%d", cur.Val)
			if i < this.Size-1 {
				result += " -> "
			}
			cur = *cur.Next
			i++
		}
	}
	return result
}
