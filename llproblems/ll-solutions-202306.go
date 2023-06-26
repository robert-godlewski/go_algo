// Solutions to linked list problems in June 2023
package llproblems

import (
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

// Merge Two Sorted Lists - Need to study this
func mergeTwoLists(list1 *linkedlists.ListNode, list2 *linkedlists.ListNode) *linkedlists.ListNode {
	if list1 == nil && list2 == nil {
		return nil
	} else if list1 == nil && list2 != nil {
		return list2
	} else if list1 != nil && list2 == nil {
		return list1
	} else {
		var temp *linkedlists.ListNode = nil
		if list1.Val <= list2.Val {
			temp = list1.Next
			list2 = mergeTwoLists(temp, list2)
			list1.Next = list2
			return list1
		} else {
			temp = list2.Next
			list1 = mergeTwoLists(list1, temp)
			list2.Next = list1
			return list2
		}
	}
}

// Add Two Numbers
func addTwoNumbers(l1 *linkedlists.ListNode, l2 *linkedlists.ListNode) *linkedlists.ListNode {
	resultList := new(linkedlists.LinkedList)
	cur1 := l1
	cur2 := l2
	addRemainder := false
	for cur1 != nil && cur2 != nil {
		sum := 0
		sum += cur1.Val
		sum += cur2.Val
		if addRemainder == true {
			sum += 1
			addRemainder = false
		}
		if sum >= 10 {
			sum -= 10
			addRemainder = true
		}
		resultList.AddAtTail(sum)
		cur1 = cur1.Next
		cur2 = cur2.Next
	}
	// if cur1 != nil but cur2 == nil
	for cur1 != nil {
		sum := 0
		sum += cur1.Val
		if addRemainder == true {
			sum += 1
			addRemainder = false
		}
		if sum >= 10 {
			sum -= 10
			addRemainder = true
		}
		resultList.AddAtTail(sum)
		cur1 = cur1.Next
	}
	// if cur2 != nil but cur1 == nil
	for cur2 != nil {
		sum := 0
		sum += cur2.Val
		if addRemainder == true {
			sum += 1
			addRemainder = false
		}
		if sum >= 10 {
			sum -= 10
			addRemainder = true
		}
		resultList.AddAtTail(sum)
		cur2 = cur2.Next
	}
	if addRemainder == true {
		resultList.AddAtTail(1)
	}
	return resultList.Head
}
