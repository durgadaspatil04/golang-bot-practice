package main

import "fmt"

func main() {
	//various type of multiple Variable can be declared in single line
	var Fname, Age = "Dsp", 28 // using type Infernce
	name, age := "Rajhans", 32 // using Short Declaration
	fmt.Println(name, age)
	fmt.Println(Fname, Age)
}
