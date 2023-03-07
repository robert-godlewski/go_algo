// This is a group of challenge problems for leetcode in March 2023
package challenges

import (
	"fmt"
)

// Running sum of a 1d Array
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

// Fizz Buzz
func FizzBuzzPrint(n int) {
	fmt.Printf("Fizz Buzz print for length of %v\n", n)
	result := fizzBuzz(n)
	fmt.Printf("Result = %v\n", result)
}

// Solved in 5 min
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
