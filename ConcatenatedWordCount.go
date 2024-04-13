package main

import (
	"fmt"
)

// getConcatenateWordCount takes two slices of strings and returns a map of word counts
func getConcatenateWordCount(stream1, stream2 []string) map[string]int {
	wordCountMap := make(map[string]int)

	// Count words from the first stream
	for _, word := range stream1 {
		wordCountMap[word]++
	}

	// Count words from the second stream
	for _, word := range stream2 {
		wordCountMap[word]++
	}
	return wordCountMap
}

func main() {
	stream1 := []string{"cat", "rat", "dog"}
	stream2 := []string{"cat", "dog", "cow", "dock"}

	wordCountMap := getConcatenateWordCount(stream1, stream2)

	// Print the concatenated word count
	for word, count := range wordCountMap {
		fmt.Printf("%s: %d\n", word, count)
	}
}
