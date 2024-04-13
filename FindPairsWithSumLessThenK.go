package main

import (
	"fmt"
	"sort"
)

// getAllPairsWithSumLessThanK finds all pairs in list A where the sum is less than K
func getAllPairsWithSumLessThanK(A []int, K int) [][2]int {
	var pairsList [][2]int
	sort.Ints(A) // Sort the slice in ascending order
	left, right := 0, len(A)-1

	for left < right {
		currentSum := A[left] + A[right]

		if currentSum < K {
			// Add all pairs with the current left and all elements from right to left
			for i := right; i > left; i-- {
				pairsList = append(pairsList, [2]int{A[left], A[i]})
			}
			left++
		} else {
			right--
		}
	}

	return pairsList
}

func main() {
	A := []int{2, 7, 11, 15}
	K := 19

	result := getAllPairsWithSumLessThanK(A, K)

	if len(result) > 0 {
		fmt.Println("Pairs found:")
		for _, pair := range result {
			fmt.Printf("%d, %d\n", pair[0], pair[1])
		}
	} else {
		fmt.Println("No such pairs found.")
	}
}
