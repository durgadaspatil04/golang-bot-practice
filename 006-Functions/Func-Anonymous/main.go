package main

import "fmt"

//Assigning function to the variable.
var (
	area = func(a, b int) int {
		return a * b

	}
)

func main() {
	fmt.Println("area is ", area(8, 5))

	func(l, w int) {
		fmt.Println(l * w)
	}(10, 20) //Passing arguments to anonymous functions.
}
