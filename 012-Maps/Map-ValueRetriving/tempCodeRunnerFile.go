package main

import "fmt"

func main() {
	NameAge := map[string]int{
		"John":  21,
		"James": 27,
		"Rob":   34,
	}

	//try retrive element from giving key to it
	na := "John"
	NewElement := NameAge[na]
	fmt.Println("1st method - Element is", NewElement)
	fmt.Println("2nd method - Element is", NameAge["Rob"])       //direct retriving using key
	fmt.Println("If Employee Not Present Case", NameAge["Mike"]) //it give zero Value if Element not present
}
