package main

import "fmt"

func main() {
    fmt.Println(isAdamInteger(12))    // true
    fmt.Println(isAdamInteger(36))    // false
    fmt.Println(isAdamInteger(5))     // false
    fmt.Println(isAdamInteger(65432)) // false
}

func isAdamInteger(num int) bool {
    return reverseInteger(reverseInteger(num)*reverseInteger(num)) == num*num
}

