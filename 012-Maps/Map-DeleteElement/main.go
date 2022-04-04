package main

import "fmt"

func main() {
	Stud := map[string]int{
		"John":  12,
		"Bond":  11,
		"James": 14,
		"Snow":  15,
	}
	fmt.Println("Map Before Delete Func Call", Stud)
	delete(Stud, "Bond") //Delete func
	fmt.Println("Map After Delete Func Call", Stud)
}
