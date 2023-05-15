// Solutions solved in early 2023
package challenges

import (
	"fmt"
	"go_algo/linkedlists"
)

// Running sum of a 1d Array
// Solved in 5 min
// O(n) time solution
// O(1) space solution
func runningSum(nums [5]int) [5]int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// Richest Customer Wealth
// Solved in 10 min
// O(n**2) time solution
// O(1) space solution
func maximumWealth(accounts [3][3]int) int {
	m := len(accounts)
	n := len(accounts[0])
	maxSum := 0
	for i := 0; i < m; i++ {
		subTotal := 0
		for j := 0; j < n; j++ {
			subTotal += accounts[i][j]
		}
		if subTotal > maxSum {
			maxSum = subTotal
		}
	}
	return maxSum
}

// Fizz Buzz
// Solved in 5 min
// O(n) time solution
// O(1) space solution
func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		val := ""
		if i%3 == 0 {
			val += "Fizz"
		}
		if i%5 == 0 {
			val += "Buzz"
		}
		if val == "" {
			val = fmt.Sprintf("%d", i)
		}
		result = append(result, val)
	}
	return result
}

// Number of steps to reduce a number to 0
// Solved in 5 min
// O(n) time solution
// O(1) space solution
func numberOfSteps(num int) int {
	steps := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps++
	}
	return steps
}

// Middle of the Linked List
// Solved in 15 min
// Best Time = 0(1)
// Normal Time = O(n)
// Space = O(1)
func middleNode(head *linkedlists.ListNode) *linkedlists.ListNode {
	size := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		size++
	}
	result := head
	i := 0
	for result != nil && i < size/2 {
		result = result.Next
		i++
	}
	return result
}

// Ransom Note
// Solved in over 1 hr - need to work on this
// Time = O(n)
// Space = O(n)
func canConstruct(ransomNote string, magazine string) bool {
	// There are 26 letters in the alphabet
	mag_arr := make([]int, 26)
	// add to the hashmap
	for _, val := range magazine {
		//fmt.Printf("i = %v\nval = %v\nval-'a' = %v\n", i, val, val-'a')
		mag_arr[val-'a']++
	}
	//fmt.Printf("mag_arr = %v\n", mag_arr)
	for _, val := range ransomNote {
		mag_arr[val-'a']--
		if mag_arr[val-'a'] < 0 {
			return false
		}
	}
	return true
}
