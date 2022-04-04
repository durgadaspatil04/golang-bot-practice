package main

import "fmt"

type employee struct {
	name   string
	salary float32
	age    uint8
}

func main() {
	emp1 := &employee{"James Bond", 23000.0, 55}

	fmt.Println("Address of struct", (emp1).name)
	fmt.Println("Employee 1 data", (*emp1).name, (*emp1).salary, *&emp1.age)
}
