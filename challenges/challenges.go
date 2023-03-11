// This is a group of challenge problems for leetcode in March 2023
package challenges

import (
	"fmt"
	//"go_algo/LinkedLists"
)

// Running sum of a 1d Array
func RunningSumPrint(nums [5]int) {
	fmt.Printf("Initial values are %v\n", nums)
	sumList := runningSum(nums)
	fmt.Printf("The running sum is %v\n", sumList)
}

// Actually runs the code
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
func MaxWealthPrint(accounts [3][3]int) {
	fmt.Printf("Account values = %v\n", accounts)
	maxClient := maximumWealth(accounts)
	fmt.Printf("The richest client has %v\n", maxClient)
}

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
func FizzBuzzPrint(n int) {
	fmt.Printf("Fizz Buzz print for length of %v\n", n)
	result := fizzBuzz(n)
	fmt.Printf("Result = %v\n", result)
}

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
func NumberOfStepsPrint(num int) {
	result := numberOfSteps(num)
	fmt.Printf("The number of steps to take %v to 0 = %v.\n", num, result)
}

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
// func MiddleNodePrint(head *linkedlists.ListNode) {}

// func middleNode(head *linkedlists.ListNode) *linkedlists.ListNode {
//     // Bad
// 	return head
// }
