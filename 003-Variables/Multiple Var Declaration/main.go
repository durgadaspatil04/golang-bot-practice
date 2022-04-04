package main

import "fmt"

func main() {
	var age, id int = 29, 001
	// multiple variables declared in of single type in single line
	var Fname, Lname = "Durgesh", "Patil"
	// Multiple VAR Declaration and INIALISED using Type Inference
	fmt.Println("values are ", age, id)
	fmt.Println("Name is ", Fname, Lname)

}
