package main

import "fmt"

func change(val *int) {
	*val = 55
}

func main() {
	a := 77
	b := &a
	fmt.Printf("The Value of a- %d Before Func call\n", a)
	change(b) //address of "a" gela so change in value of "a" hoiel
	fmt.Printf("The Value of a- %d Before Func call\n", a)
}
