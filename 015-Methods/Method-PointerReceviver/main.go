package main

import "fmt"

type person struct {
	name   string
	age    int
	salary float32
}

//Method with value receiver
func (p person) changeName(n string) {
	p.name = n
}

/*Method with pointer receiver
pointer receivers can be used when changes made to the
 receiver inside the method should be visible to the caller. */
func (p *person) changeAge(a int) {
	p.age = a
}
func main() {
	p1 := person{"Mike", 34, 30000}

	p1.changeName("James Bond")
	fmt.Println("Person info", p1)
	(&p1).changeAge(45)
	fmt.Println("After Called pointer receiver new age ", p1)
}
