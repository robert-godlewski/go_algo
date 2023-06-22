// The main code to run all off the packages for this program
package main

import (
	"fmt"
	"go_algo/binarytrees"
	"go_algo/challenges"
	"go_algo/llproblems"
	"go_algo/queuesstacks"
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
	// Testing out the Challenge problems
	thickLine()
	fmt.Println("Challenge Problems")
	thinLine()

	// Testing the RunningSum
	challenges.RunningSumTests()
	thinLine()

	// Testing the MaximumWealth
	challenges.MaxWealthTests()
	thinLine()

	// Testing FizzBuzz
	challenges.FizzBuzzTests()
	thinLine()

	// Testing Number of Steps
	challenges.NumberOfStepsTests()
	thinLine()

	// Testing Middle of the Linked List
	challenges.MiddleNodeTests()
	thinLine()

	// Testing Ransom Note
	challenges.CanConstructTests()

	// Testing Linked Lists
	thickLine()
	fmt.Println("Running Linked List Tests")
	thickLine()

	// Testing reverseList
	llproblems.ReverseLLTests()
	thinLine()

	// Testing removeElements
	llproblems.RemoveListNodeTests()
	thinLine()

	// Testing Palindronme Linked Lists
	llproblems.PalindromeLLTests()
	thinLine()

	// Testing Merge Two Sorted Lists
	llproblems.MergeLLTests()

	thickLine()
	fmt.Println("Testing out class features")
	thickLine()

	// Testing out Queue class
	var testQ = new(queuesstacks.LinkedQueue)
	//fmt.Println("Empty linked queue:")
	//var testQStr = testQ.PrintQueue()
	//fmt.Printf("%v\n", testQStr)
	//fmt.Println("Queue list add in operations:")
	testQ.Push(1)
	//testQStr = testQ.PrintQueue()
	//fmt.Printf("%v\n", testQStr)
	testQ.Push(2)
	//testQStr = testQ.PrintQueue()
	//fmt.Printf("%v\n", testQStr)
	//fmt.Println("Queue list removal operation:")
	//q_val :=
	testQ.Pop()
	//fmt.Printf("Removed %v from the queue\n", q_val)
	//testQStr = testQ.PrintQueue()
	//fmt.Printf("%v\n", testQStr)
	thinLine()

	// Testing out Stack class
	var testS = new(queuesstacks.LinkedStack)
	//fmt.Println("Empty linked stack:")
	//var testSstr = testS.PrintStack()
	//fmt.Printf("%v\n", testSstr)
	//fmt.Println("Stack list add in operations:")
	testS.Push(1)
	//testSstr = testS.PrintStack()
	//fmt.Printf("%v\n", testSstr)
	testS.Push(2)
	//testSstr = testS.PrintStack()
	//fmt.Printf("%v\n", testSstr)
	//fmt.Println("Stack list removal operation:")
	//s_val :=
	testS.Pop()
	//fmt.Printf("Removed %v from the stack\n", s_val)
	//testSstr = testS.PrintStack()
	//fmt.Printf("%v\n", testSstr)
	thinLine()

	// Testing out Binary Tree Package
	testBT := new(binarytrees.BinaryTree)
	testBT.Push(20)
	testBT.Push(10)
	testBT.Push(30)
	fmt.Printf("New binary tree node with value of %v\n", testBT.Root.Val)
	fmt.Printf("Binary Tree left node = %v\n", testBT.Root.Left.Val)
	fmt.Printf("Binary Tree right node = %v\n", testBT.Root.Right.Val)
	testBST1 := testBT.BST(5)
	fmt.Printf("Able to find 5? %v\n", testBST1)
	testBST2 := testBT.BST(10)
	fmt.Printf("Able to find 10? %v\n", testBST2)
}
