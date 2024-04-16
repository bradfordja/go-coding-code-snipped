package main

import (
	"fmt"
	"sort"
	"strings"
)

// groupAnagrams groups strings such that each group contains strings that are anagrams of each other
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	anagramMap := make(map[string][]string)
	for _, s := range strs {
		// Sort the characters of the string to form the key
		sorted := sortString(s)
		anagramMap[sorted] = append(anagramMap[sorted], s)
	}

	// Convert the map values to a slice of slices
	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

// sortString sorts the characters of a string and returns a new sorted string
func sortString(s string) string {
	// Convert string to a slice of runes to handle Unicode correctly
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	groupedAnagrams := groupAnagrams(strs)
	fmt.Println("Grouped Anagrams:", groupedAnagrams)
}
