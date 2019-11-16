package main

import (
	"bufio"
	"fmt"
	"os"
)

// Scan Stdin for an integer to use as counter and return it.
func scanForCounter(reader *bufio.Reader) int {
	var input int

	// Scan Stdin for the next integer
	fmt.Fscan(reader, &input)

	// Validate the integer value
	if input <= 0 || input > 100 {
		return 0
	}

	return input
}

// Scan the next given number of integers and calculate the sum of their square values.
func calculateSumOfSquares(reader *bufio.Reader, integersCount int) int {
	// When we reach the desired number of integers, flush the rest of the input and return
	if integersCount <= 0 {
		return 0
	}

	var input int

	// Scan Stdin for the next integer of the test case
	fmt.Fscan(reader, &input)

	// Prepare the increment to add to the sum result
	var increment int

	// Ensure that we set the increment to the square value of the integer
	// only if it's not negative and not greater than 100
	if input > 0 && input <= 100 {
		increment = input * input
	}

	// Recursively add the square value of each valid integer until we reach the desired number of integers processed
	return increment + calculateSumOfSquares(reader, integersCount-1)
}

// Process a test case.
func processTestCase(reader *bufio.Reader) {
	// Scan Stdin for the number of integers the test case will contain
	integersCount := scanForCounter(reader)

	// Get the sum of the square values of positive integers
	result := calculateSumOfSquares(reader, integersCount)

	// Display the result
	fmt.Println(result)
}

// Process the given number of test cases.
func processTestCases(reader *bufio.Reader, testCasesCount int) {
	if testCasesCount <= 0 {
		return
	}

	// Process the next test case
	processTestCase(reader)

	// Recursive call to process the next test case
	processTestCases(reader, testCasesCount-1)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Don't forget to clear the reader's buffer once we're done
	defer reader.Discard(reader.Buffered())

	testCasesCount := scanForCounter(reader)

	processTestCases(reader, testCasesCount)
}
