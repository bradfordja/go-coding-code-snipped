package main

import (
	"fmt"
)

// canPermutePalindrome checks if any permutation of the string can be a palindrome
func canPermutePalindrome(s string) bool {
	// Map to store character counts
	charCounts := make(map[rune]int)

	// Count each character
	for _, c := range s {
		charCounts[c]++
	}

	// Count the number of characters with an odd count
	oddCount := 0
	for _, count := range charCounts {
		if count%2 != 0 {
			oddCount++
		}

		// If more than one character has an odd count, return false early
		if oddCount > 1 {
			return false
		}
	}

	// A string can form a palindrome if there is at most one odd count
	return true
}

func main() {
	input1 := "tactcoa"
	input2 := "java"

	fmt.Printf("Can \"%s\" permute to a palindrome? %v\n", input1, canPermutePalindrome(input1))
	fmt.Printf("Can \"%s\" permute to a palindrome? %v\n", input2, canPermutePalindrome(input2))
}
