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
}

func reverseLLPrint(list *linkedlists.LinkedList) {
	fmt.Println("List in:")
	fmt.Printf("%v\n", list.PrintLL())
	result := reverseList(list.Head)
	fmt.Printf("The head of the result = (%v)\n", result.Val)
}

// Removed Linked List
func RemoveListNodeTests() {
	sll1 := new(linkedlists.LinkedList)
	sll1.AddAtHead(1)
	sll1.AddAtTail(2)
	sll1.AddAtTail(6)
	sll1.AddAtTail(3)
	sll1.AddAtTail(4)
	sll1.AddAtTail(5)
	sll1.AddAtTail(6)
	removeLNPrint(sll1, 6)
	sll2 := new(linkedlists.LinkedList)
	sll2.AddAtHead(7)
	sll2.AddAtTail(7)
	sll2.AddAtTail(7)
	sll2.AddAtTail(7)
	removeLNPrint(sll2, 7)
}

func removeLNPrint(list *linkedlists.LinkedList, val int) {
	fmt.Println("List in:")
	fmt.Printf("%v\n", list.PrintLL())
	result := removeElements(list.Head, val)
	fmt.Println("Result:")
	cur := result
	resultStr := "(head) -> "
	for cur != nil {
		resultStr += fmt.Sprintf("%d", cur.Val)
		resultStr += " -> "
		cur = cur.Next
	}
	resultStr += "(tail)"
	fmt.Printf("%v\n", resultStr)
	//fmt.Printf("The head of the result = (%v)\n", result.Val)
}

// Palindrome Linked List
func PalindromeLLTests() {
	sll1 := new(linkedlists.LinkedList)
	sll1.AddAtTail(1)
	sll1.AddAtTail(2)
	sll1.AddAtTail(2)
	sll1.AddAtTail(1)
	palindromeLLPrint(sll1)
	sll2 := new(linkedlists.LinkedList)
	sll2.AddAtTail(1)
	sll2.AddAtTail(2)
	palindromeLLPrint(sll2)
}

func palindromeLLPrint(list *linkedlists.LinkedList) {
	fmt.Println("Is this list a Palindrome?")
	fmt.Printf("%v\n", list.PrintLL())
	is_palindrome := isPalindrome(list.Head)
	fmt.Printf("%v\n", is_palindrome)
}
