package main

import "fmt"

type person struct {
	name string
	age  int
}

//Info method has person as the receiver type
func (p person) Info() { //method declaration and defination of struct type of receiver
	fmt.Println("Person info is", p.name, p.age)
}
func main() {
	p1 := person{"James", 54}
	p1.Info()

}
