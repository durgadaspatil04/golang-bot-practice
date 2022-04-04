package main

import "fmt"

func main() {
	const st = "Hello World" //untyped String Const
	type new string
	const a new = "hello"
	fmt.Printf("type is %T value %v \n", a, a)
	fmt.Printf("type is %T value %v \n", st, st)
	const str string = "India" //Typed Constant
	fmt.Printf("type is %T value %v ", str, str)

}
