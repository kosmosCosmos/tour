package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	firstFibonacciIndex, secondFibonacciIndex := 0, 1

	fiboFunc := func() int {
		firstFibonacciIndex, secondFibonacciIndex = secondFibonacciIndex, firstFibonacciIndex+secondFibonacciIndex

		return secondFibonacciIndex
	}

	return fiboFunc
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
