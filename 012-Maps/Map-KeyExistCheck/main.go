package main

import "fmt"

func main() {
	Emp := map[string]int{
		"Mike": 1000,
		"Bond": 2000,
	}
	//syntax to check key present or not
	check := "Shaktiman"
	// ok is Boolean returns true-Present flase-Not Present value Returns Value OF the key
	if value, ok := Emp[check]; ok{
		fmt.Println("The Salary of Mike", value)
	}
	fmt.Println("The Key NOt Found!!")
	fmt.Printf("Type ok - %T Value %v\n", ok, value)
}
