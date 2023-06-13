package llproblems

import (
	"fmt"
	"go_algo/linkedlists"
)

// Reverse Linked List
func ReverseLLTests() {
	sll1 := new(linkedlists.LinkedList)
	sll1.AddAtHead(1)
	sll1.AddAtTail(2)
	sll1.AddAtTail(3)
	sll1.AddAtTail(4)
	sll1.AddAtTail(5)
	reverseLLPrint(sll1)
	sll2 := new(linkedlists.LinkedList)
	sll2.AddAtHead(1)
	sll2.AddAtTail(2)
	reverseLLPrint(sll2)
	reverseLLPrint(nil)
}

func reverseLLPrint(list *linkedlists.LinkedList) {
	fmt.Println("List in:")
	fmt.Printf("%v\n", list.PrintLL())
	result := reverseList(list.Head)
	fmt.Printf("The head of the result = (%v)\n", result.Val)
}
