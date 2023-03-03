// The main code to run all off the packages for this program
package main

import (
	"fmt"
	"go_algo/challenges"
)

// Use to make a thickLine in the terminal to split up the main categories
func thickLine() {
	fmt.Println("==============")
}

// Use to make a thin line in the terminal to split up the code tests
func thinLine() {
	fmt.Println("--------------")
}

// Testing out Challenge Problems

func main() {
	// Testing out the input
	// var name string
	// var numTemp int
	// fmt.Println("Enter your name:")
	// fmt.Scan(&name)
	// fmt.Println("Enter in a number")
	// fmt.Scan(&numTemp)
	// fmt.Printf("Hello %v, you chose %v\n", name, numTemp)

	// Testing out the Challenge problems
	thickLine()
	fmt.Println("Challenge Problems")
	thinLine()

	// Testing the RunningSum
	fmt.Println("Checking the Running Sum of 1d Arrays")
	numSum := [5]int{1, 2, 3, 4}
	challenges.RunningSumPrint(numSum)
	numSum[1] = 1
	numSum[2] = 1
	numSum[3] = 1
	numSum[4] = 1
	challenges.RunningSumPrint(numSum)
	numSum[0] = 3
	numSum[1] = 1
	numSum[2] = 2
	numSum[3] = 10
	numSum[4] = 1
	challenges.RunningSumPrint(numSum)
}
