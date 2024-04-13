package main

import (
	"fmt"
)

// analyze processes an array of numbers and identifies missing and duplicated numbers within a range.
func analyze(nums []int) map[string][]int {
	result := make(map[string][]int)

	// This slice helps track the count of numbers, array size of 1000001 (to include the number 1000000).
	tracker := make([]int, 1000001)

	// Iterate through the nums and update our tracker slice.
	for _, num := range nums {
		if num > 0 && num <= 1000000 {
			tracker[num]++
		}
	}

	var missing []int
	var duplicated []int

	// Now, process our tracker slice.
	for i := 1; i <= 1000000; i++ {
		if tracker[i] == 0 {
			missing = append(missing, i)
		} else if tracker[i] > 1 {
			duplicated = append(duplicated, i)
		}
	}

	result["missing"] = missing
	result["duplicated"] = duplicated

	return result
}

func main() {
	numArray := []int{-10, -4, -4, 0, 1, 3, 3, 3, 3, 5, 6, 6, 999999, 1000000000, 1000000000, 1000000001}
	result := analyze(numArray)

	fmt.Println("Missing:", result["missing"])
	fmt.Println("Duplicated:", result["duplicated"])
}
