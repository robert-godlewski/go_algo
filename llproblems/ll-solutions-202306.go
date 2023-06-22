// Solutions to linked list problems in June 2023
package llproblems

import (
	"fmt"
	"go_algo/linkedlists"
)

// Reverse Linked List Solution
// Time Complexity = O(1)
// Space Complexity = O(n)
// Solved in 1 hr
func reverseList(head *linkedlists.ListNode) *linkedlists.ListNode {
	return reverseListSolution(head, nil)
}

func reverseListSolution(head *linkedlists.ListNode, prev *linkedlists.ListNode) *linkedlists.ListNode {
	if head != nil {
		cur := head
		var next *linkedlists.ListNode = nil
		if head.Next != nil {
			next = reverseListSolution(head.Next, cur)
		}
		cur.Next = prev
		if prev != nil {
			prev.Next = next
		}
		if next != nil {
			head = next
		}
	}
	return head
}

// Remove Linked List Solution
func removeElements(head *linkedlists.ListNode, val int) *linkedlists.ListNode {
	if head != nil {
		cur := head
		var prev *linkedlists.ListNode = nil
		for cur != nil {
			if cur.Val == val {
				next := cur.Next
				if prev != nil {
					prev.Next = next
				}
				if next != nil {
					next.Prev = prev
				}
				if cur == head {
					head = next
				}
			} else {
				prev = cur
			}
			cur = cur.Next
		}
	}
	return head
}

// Palindrome Linked List
func isPalindrome(head *linkedlists.ListNode) bool {
	if head != nil {
		var nodeList [100000]int
		cur := head
		i := 0
		for cur != nil && i < len(nodeList) {
			nodeList[i] = cur.Val
			cur = cur.Next
			i++
		}
		i--
		j := 0
		for j < i {
			if nodeList[i] != nodeList[j] {
				return false
			}
			j++
			i--
		}
		return true
	}
	return false
}

// Merge Two Sorted Lists
// Bad solution according to leetcode
func mergeTwoLists(list1 *linkedlists.ListNode, list2 *linkedlists.ListNode) *linkedlists.ListNode {
	if list2 != nil && list1 == nil {
		return list2
	} else {
		var prev *linkedlists.ListNode = nil
		cur1 := list1
		cur2 := list2
		for cur1 != nil && cur2 != nil {
			fmt.Printf("cur1 = %v\n", cur1.Val)
			fmt.Printf("cur2 = %v\n", cur2.Val)
			if cur1.Val > cur2.Val {
				fmt.Printf("%v > %v\n", cur1.Val, cur2.Val)
				prev.Next = cur2
				temp := cur2.Next
				cur2.Next = cur1
				cur2 = temp
			} else {
				fmt.Printf("%v <= %v\n", cur1.Val, cur2.Val)
				temp := cur1.Next
				cur1.Next = cur2
				prev = cur1
				cur1 = cur1.Next
				cur2 = temp
			}
		}
		return list1
	}
}
