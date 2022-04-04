package main

import "fmt"

func change(n [5]int) {
	n[0] = 20
	fmt.Println("inside Function", n)
}
func main() {
	n := [5]int{1, 2, 3, 4, 5}
	fmt.Println("array before Func call", n)
	change(n) //Bcoz Array is pass by value so Changes doesn't refelct what r made in function
	fmt.Println("array after Func call", n)
}
