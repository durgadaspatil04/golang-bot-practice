package main

import "fmt"

func main() {
	var employeeSalary map[string]int
	/* " employeeSalary["steve"] = 12000 " if we try this way to initialise it give error
	Bcoz map is nil so can"t initialise this way*/

	//so you have use this method to initiliased
	employeeSalary = map[string]int{
		"hello": 20,
	}
	fmt.Println("First Map", employeeSalary)

	emp := map[string]int{
		"John":     10,
		"Jonny":    20,
		"Janardan": 30,
	}
	emp["TarRamPam"] = 2000
	fmt.Println("Initialised", emp)
	//iterating map element using range
	for key, value := range emp {
		fmt.Println("Using Range", key, value)
	}
}
