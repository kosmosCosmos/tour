package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	WalkHelper(t, ch)
	close(ch)
}

func WalkHelper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	// Visit all left sub-children
	if t.Left != nil {
		WalkHelper(t.Left, ch)
	}

	// Send curent value to channel
	ch <- t.Value

	// Visit all right sub-children
	if t.Right != nil {
		WalkHelper(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)
	ch2 := make(chan int)
	go Walk(t2, ch2)

	for {
		value1, ok1 := <-ch1
		value2, ok2 := <-ch2
		if value1 != value2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	ch := make(chan int)

	fmt.Println("Sample Tree Traversal:")
	go Walk(tree.New(1), ch)

	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
