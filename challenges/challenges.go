// This is a group of challenge problems for leetcode in March 2023
package challenges

import "fmt"

// Running sum of a 1d Array
// Printing it out
func RunningSumPrint(nums [5]int) {
	fmt.Printf("Initial values are %v\n", nums)
	sumList := runningSum(nums)
	fmt.Printf("The running sum is %v\n", sumList)
}

// Actually runs the code
// Solved in 5 min
func runningSum(nums [5]int) [5]int {
	for i := 1; i < len(nums); i++ {
		nums[i] += nums[i-1]
	}
	return nums
}

// Richest Customer Wealth
// Printing it out
func MaxWealthPrint(accounts [3][3]int) {
	fmt.Printf("Account values = %v\n", accounts)
	maxClient := maximumWealth(accounts)
	fmt.Printf("The richest client has %v\n", maxClient)
}

// Solved in 12pm
// Solved in 10 min
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
