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
	sl := []int{10, 20, 30, 40, 50}
	Find(20, sl...) //if you only write the sl it will give you error
}
