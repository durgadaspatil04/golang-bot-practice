package main

import "fmt"

func main() {
	n := 5
	for i := 0; i < n; i++ { //nested loop, loop inside loop
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}
