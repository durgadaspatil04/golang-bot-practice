package main

import "fmt"

func main() {
	var a = [5]int{10, 20, 30, 40, 50}
	var b []int = a[1:5] //create a slice of a[1] to a[3] slice from array
	fmt.Println(b)
	//slice are reference type
	c := []int{10, 2, 34, 5, 6, 77, 8} //slice short hand declaration
	fmt.Println(c)

}
