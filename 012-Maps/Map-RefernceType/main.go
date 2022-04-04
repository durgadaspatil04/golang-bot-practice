package main

import "fmt"

func main() {
	employeInfo := map[string]int{
		"John":  1500,
		"James": 1800,
		"Mike":  5000,
	}
	fmt.Println("Orginal", employeInfo)
	Modified := employeInfo //assigned to new variable but both points to same location
	Modified["Mike"] = 2500 //this change also reflect in Orginal map too
	fmt.Println("Modified", Modified)
	fmt.Println("Orginal got changed", employeInfo)
}
