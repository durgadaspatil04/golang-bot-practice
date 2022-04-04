//It is possible to declare structs without creating a new data type. These types
// of structs are called anonymous structs.
package main

import "fmt"

func main() {
	emp1 := struct {
		name   string
		age    int
		salary float64
	}{"Dsp", 23, 234.00}
	emp2 := struct {
		name   string
		age    int
		salary float64
	}{
		name:   "James",
		age:    54,
		salary: 45004.00}
	fmt.Println("Anonymous Structs Are", emp1, "\n", emp2)
}
