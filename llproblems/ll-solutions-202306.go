// Solutions to linked list problems in June 2023
package llproblems

import "go_algo/linkedlists"

// Reverse Linked List Solution
// Time Complexity =
// Space Complexity =
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
