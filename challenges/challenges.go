// This is a group of challenge problems for leetcode in March 2023
package challenges

import (
	"fmt"
	"go_algo/linkedlists"
)

// Running sum of a 1d Array
func RunningSumTests() {
	fmt.Println("Checking the Running Sum of 1d Arrays")
	numSum := [5]int{1, 2, 3, 4}
	runningSumPrint(numSum)
	numSum[1] = 1
	numSum[2] = 1
	numSum[3] = 1
	numSum[4] = 1
	runningSumPrint(numSum)
	numSum[0] = 3
	numSum[1] = 1
	numSum[2] = 2
	numSum[3] = 10
	numSum[4] = 1
	runningSumPrint(numSum)
}

func runningSumPrint(nums [5]int) {
	fmt.Printf("Initial values are %v\n", nums)
	sumList := runningSum(nums)
	fmt.Printf("The running sum is %v\n", sumList)
}

// Richest Customer Wealth
func MaxWealthTests() {
	fmt.Println("Finding the Maximum Wealth")
	accounts1 := [3][3]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	maxWealthPrint(accounts1)
	accounts2 := [3][3]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	maxWealthPrint(accounts2)
	accounts3 := [3][3]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}
	maxWealthPrint(accounts3)
}

func maxWealthPrint(accounts [3][3]int) {
	fmt.Printf("Account values = %v\n", accounts)
	maxClient := maximumWealth(accounts)
	fmt.Printf("The richest client has %v\n", maxClient)
}

// Fizz Buzz
func FizzBuzzTests() {
	fmt.Println("Fizz Buzz")
	fizzBuzzPrint(3)
	fizzBuzzPrint(5)
	fizzBuzzPrint(15)
}

func fizzBuzzPrint(n int) {
	fmt.Printf("Fizz Buzz print for length of %v\n", n)
	result := fizzBuzz(n)
	fmt.Printf("Result = %v\n", result)
}

// Number of steps to reduce a number to 0
func NumberOfStepsTests() {
	fmt.Println("Number of Steps")
	numberOfStepsPrint(14)
	numberOfStepsPrint(8)
	numberOfStepsPrint(123)
}

func numberOfStepsPrint(num int) {
	result := numberOfSteps(num)
	fmt.Printf("The number of steps to take %v to 0 = %v.\n", num, result)
}

// Middle of the Linked List
func MiddleNodeTests() {
	sllt := new(linkedlists.LinkedList)
	sllt.AddAtHead(1)
	sllt.AddAtTail(2)
	sllt.AddAtTail(3)
	sllt.AddAtTail(4)
	sllt.AddAtTail(5)
	middleNodePrint(sllt.Head)
	sllt.AddAtTail(6)
	middleNodePrint(sllt.Head)
}

func middleNodePrint(head *linkedlists.ListNode) {
	result := middleNode(head)
	fmt.Printf("The head of the result = (%v)\n", result.Val)
}

// Ransom Note
func CanConstructTests() {
	canConstructPrint("a", "b")
	canConstructPrint("aa", "ab")
	canConstructPrint("aa", "aab")
}

func canConstructPrint(ransomNote string, magazine string) {
	result := canConstruct(ransomNote, magazine)
	fmt.Printf("Can we use '%v' with '%v'? %v\n", ransomNote, magazine, result)
}
