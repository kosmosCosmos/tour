package main

import (
	"code.google.com/p/go-tour/wc"
	"strings"
)

// Implement WordCount.
// It should return a map of the counts of each “word” in the string s.
// The wc.Test function runs a test suite against the provided function and prints success or failure.
// You might find strings.Fields helpful.
func WordCount(s string) map[string]int {

	wordCountMap := make(map[string]int)
	arrayOfWords := strings.Fields(s)

	for _, word := range arrayOfWords {
		wordCountMap[word] = wordCountMap[word] + 1
	}

	return wordCountMap
}

func main() {
	wc.Test(WordCount)
}
