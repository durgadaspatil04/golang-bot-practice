package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println("Value of a is", a)
	fmt.Println("Value of b is", b)
	c := a && b
	fmt.Println("Value of C is", c)
	d := a || b
	fmt.Println("Value of D is", d)
}
