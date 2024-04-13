package main

import (
	"fmt"
)

// factorial computes the factorial of n using recursion
func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	num := 7
	result := factorial(num)
	fmt.Printf("Factorial of %d is: %d\n", num, result)
}

