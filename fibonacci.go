package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	firstFibonacciIndex := 0
	secondFibonacciIndex := 1

	fiboFunc := func() int {
		returnValue := firstFibonacciIndex + secondFibonacciIndex
		firstFibonacciIndex = secondFibonacciIndex
		secondFibonacciIndex = returnValue

		return returnValue
	}

	return fiboFunc
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
