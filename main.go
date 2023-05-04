// The main code to run all off the packages for this program
package main

import (
	"fmt"
	"go_algo/challenges"
	"go_algo/linkedlists"
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
	thinLine()

	// Testing the MaximumWealth
	fmt.Println("Finding the Maximum Wealth")
	accounts1 := [3][3]int{
		{1, 2, 3},
		{3, 2, 1},
	}
	challenges.MaxWealthPrint(accounts1)
	accounts2 := [3][3]int{
		{1, 5},
		{7, 3},
		{3, 5},
	}
	challenges.MaxWealthPrint(accounts2)
	accounts3 := [3][3]int{
		{2, 8, 7},
		{7, 1, 3},
		{1, 9, 5},
	}
	challenges.MaxWealthPrint(accounts3)
	thinLine()

	// Testing FizzBuzz
	fmt.Println("Fizz Buzz")
	challenges.FizzBuzzPrint(3)
	challenges.FizzBuzzPrint(5)
	challenges.FizzBuzzPrint(15)
	thinLine()

	// Testing Number of Steps
	fmt.Println("Number of Steps")
	challenges.NumberOfStepsPrint(14)
	challenges.NumberOfStepsPrint(8)
	challenges.NumberOfStepsPrint(123)
	thinLine()

	thickLine()
	fmt.Println("Testing out class features")
	thickLine()

	// Testing out ListNode class
	// var n1 = new(linkedlists.ListNode)
	// n1.Val = 1
	// var n2 = new(linkedlists.ListNode)
	// n2.Val = 2
	// n2.Prev = n1
	// n1.Next = n2
	// fmt.Printf("n1 val = %v\nn1.next = %v\n", n1.Val, n1.Next.Val)
	// fmt.Printf("n2 val = %v\nn2.prev = %v\n", n2.Val, n2.Prev.Val)

	// Testing out LinkedList class
	var testLL = new(linkedlists.LinkedList)
	fmt.Println("Empty linked list:")
	var testLLStr = testLL.PrintLL()
	fmt.Printf("%v\n", testLLStr)
	fmt.Println("Linked list after several add in operations:")
	testLL.AddAtHead(1)
	testLLStr = testLL.PrintLL()
	fmt.Printf("%v\n", testLLStr)
	testLL.AddAtTail(3)
	testLLStr = testLL.PrintLL()
	fmt.Printf("%v\n", testLLStr)
	testLL.AddAtIndex(1, 2)
	testLLStr = testLL.PrintLL()
	fmt.Printf("%v\n", testLLStr)
	indNum := 1
	var testInx = testLL.Get(indNum)
	fmt.Printf("Index %v = %v\n", indNum, testInx)
	testLL.DeleteAtIndex(indNum)
	testLLStr = testLL.PrintLL()
	fmt.Printf("%v\n", testLLStr)
	thinLine()

	// Testing out Queue class
	var testQ = new(queuesstacks.LinkedQueue)
	fmt.Println("Empty linked queue:")
	var testQStr = testQ.PrintQueue()
	fmt.Printf("%v\n", testQStr)
	fmt.Println("Queue list add in operations:")
	testQ.Push(1)
	testQStr = testQ.PrintQueue()
	fmt.Printf("%v\n", testQStr)
	testQ.Push(2)
	testQStr = testQ.PrintQueue()
	fmt.Printf("%v\n", testQStr)
	fmt.Println("Queue list removal operation:")
	q_val := testQ.Pop()
	fmt.Printf("Removed %v from the queue\n", q_val)
	testQStr = testQ.PrintQueue()
	fmt.Printf("%v\n", testQStr)
	thinLine()

	// Testing out Stack class
	var testS = new(queuesstacks.LinkedStack)
	fmt.Println("Empty linked stack:")
	var testSstr = testS.PrintStack()
	fmt.Printf("%v\n", testSstr)
	fmt.Println("Stack list add in operations:")
	testS.Push(1)
	testSstr = testS.PrintStack()
	fmt.Printf("%v\n", testSstr)
	testS.Push(2)
	testSstr = testS.PrintStack()
	fmt.Printf("%v\n", testSstr)
	fmt.Println("Stack list removal operation:")
	s_val := testS.Pop()
	fmt.Printf("Removed %v from the stack\n", s_val)
	testSstr = testS.PrintStack()
	fmt.Printf("%v\n", testSstr)
	thinLine()
}
