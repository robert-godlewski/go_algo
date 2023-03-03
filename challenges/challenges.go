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
