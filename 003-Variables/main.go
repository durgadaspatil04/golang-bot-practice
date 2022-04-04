package main

import (
	"fmt"
)

func main() {
	var age int      //Declaration of variable inialization is optional here when u used VAR keyword
	var name = "Dsp" /*to remove the declaration of data type while declaring
	a variable. This is generally termed as the TYPE INFERENCE.*/
	fmt.Println("the age is ", age)
	fmt.Println("the age is ", name)
	age = 10 //inialization of variable
	fmt.Println("the age is ", age)
}
