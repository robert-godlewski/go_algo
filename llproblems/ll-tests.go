package llproblems

import (
	"fmt"
	"go_algo/linkedlists"
)

// For printing out the linked list
func PrintLL(head *linkedlists.ListNode) string {
	result := "(head) -> "
	if head != nil {
		i := 0
		cur := head
		for cur != nil {
			result += fmt.Sprintf("%d", cur.Val)
			result += " -> "
			cur = cur.Next
			i++
		}
	}
	result += "(tail)"
	return result
}

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
	initial := PrintLL(list.Head)
	fmt.Printf("%v\n", initial)
	fmt.Println("Reversed list:")
	result := reverseList(list.Head)
	fmt.Printf("%v\n", PrintLL(result))
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
	fmt.Printf("%v\n", PrintLL(list.Head))
	result := removeElements(list.Head, val)
	fmt.Println("Result:")
	fmt.Printf("%v\n", PrintLL(result))
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
	fmt.Printf("%v\n", PrintLL(list.Head))
	is_palindrome := isPalindrome(list.Head)
	fmt.Printf("%v\n", is_palindrome)
}

// mergeTwoLists
func MergeLLTests() {
	sll1 := new(linkedlists.LinkedList)
	sll1.AddAtTail(1)
	sll1.AddAtTail(2)
	sll1.AddAtTail(4)
	sll2 := new(linkedlists.LinkedList)
	sll2.AddAtTail(1)
	sll2.AddAtTail(3)
	sll2.AddAtTail(4)
	mergeLLPrint(sll1, sll2)
}

func mergeLLPrint(list1 *linkedlists.LinkedList, list2 *linkedlists.LinkedList) {
	fmt.Println("Merging these lists:")
	fmt.Printf("list 1: %v\n", PrintLL(list1.Head))
	fmt.Printf("list 2: %v\n", PrintLL(list2.Head))
	merged := mergeTwoLists(list1.Head, list2.Head)
	fmt.Println("Result:")
	fmt.Printf("%v\n", PrintLL(merged))
}
