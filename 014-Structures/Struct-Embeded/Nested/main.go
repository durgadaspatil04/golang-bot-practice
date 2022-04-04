package main

import "fmt"

type address struct {
	city, country string
}
type person struct {
	name string
	age  int
	address
}

func main() {
	p1 := person{
		name: "Dsp",
		age:  23,
		address: address{
			city:    "Pune",
			country: "India",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.address, "\n", p1.city, p1.name, p1.age, p1.country)
	fmt.Println(p1.city)    //city is promoted field from address to person struct
	fmt.Println(p1.country) //country is promoted field

}
