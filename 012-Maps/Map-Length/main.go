package main

import "fmt"

func main() {
	employeSalary := map[string]int{
		"james": 2000,
		"bond":  2500,
	}
	fmt.Println("the Length of map", len(employeSalary))
	employeSalary["john"] = 4000
	fmt.Println("the Length of map after insertion", len(employeSalary))
}
