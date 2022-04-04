package main

import "fmt"

type Details interface {
	describe()
}

type address struct {
	city, country string
}

type person struct {
	name string
	age  uint8
}

func (a *address) describe() { //implemented using pointer receiver
	fmt.Println("The Address is ", a.city, a.country)
	a.city = "Pune"
}

func (p person) describe() {
	fmt.Println("Person Details is", p.name, p.age)
	p.name = "Ramesh"
}
func main() {
	var d Details
	a1 := address{"Jalgaon", "India"}
	d = &a1 //pointer recevier
	d.describe()
	p1 := person{"James", 54}
	d = p1
	d.describe()
	fmt.Println("The Address is ", a1.city, a1.country)
	//city change made in function also reflect in orginal variable
	fmt.Println("The Person info is ", p1.name, p1.age)
	//didn't change reflect value Value Recevier type
}
