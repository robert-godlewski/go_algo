// Solutions to linked list problems in June 2023
package llproblems

import "go_algo/linkedlists"

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
