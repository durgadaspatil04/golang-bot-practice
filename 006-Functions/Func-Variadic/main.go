package main

import (
	"fmt"
)

func Find(n int, nums ...int) { // this veriadic function takes 0 to variable no of argument
	fmt.Printf("Type of nums %T\n", nums)
	found := false
	for i, v := range nums {
		if n == v {
			fmt.Println(n, "The value Found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(n, "not found in ", nums)
	}
	fmt.Printf("\n")
}

func main() {
	Find(89, 12, 45, 67, 89, 56, 43)
	Find(22, 1, 2, 3, 4, 5, 22)
	Find(50, 10, 20, 30, 04, 5, 22)
	Find(1)
}
