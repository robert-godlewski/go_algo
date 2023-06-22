package linkedlists

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
	cur := *this.Head
	for i := 0; i < this.Size; i++ {
		if i == index {
			return cur.Val
		}
		cur = *cur.Next
	}
	return -1
}

func (this *LinkedList) AddAtHead(val int) {
	n := &ListNode{
		Val: val,
	}
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
	n := &ListNode{
		Val: val,
	}
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

func (this *LinkedList) AddAtIndex(index, val int) {
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.Size {
		this.AddAtTail(val)
	} else if index < this.Size {
		n := &ListNode{
			Val: val,
		}
		cur := this.Head
		prev := cur
		i := 0
		for cur != nil {
			if i == index {
				n.Prev = prev
				prev.Next = n
				n.Next = cur
				cur.Prev = n
				break
			}
			i++
			prev = cur
			cur = cur.Next
		}
		this.Size++
	}
}

func (this *LinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.Head = this.Head.Next
		this.Head.Prev = nil
	} else if index == this.Size {
		this.Tail = this.Tail.Prev
		this.Tail.Next = nil
	} else if index < this.Size {
		cur := this.Head
		prev := cur
		i := 0
		for cur != nil {
			if i == index {
				prev.Next = cur.Next
				if cur.Next != nil {
					cur.Next.Prev = prev
				}
				break
			}
			prev = cur
			cur = cur.Next
			i++
		}
	}
	if index <= this.Size {
		this.Size--
	}
}
