package main

import "fmt"

func main() {
	emp := map[string]int{
		"Mike":  20,
		"James": 22,
		"Rob":   13,
	}
	stud := map[string]int{
		"Mike":  20,
		"James": 22,
		"Rob":   13,
		// "Yash":  25,
	}
	// if emp == stud {    // u can't compare two maps using == operator
	// 	fmt.Println("Equal")
	// }
	if len(emp) == len(stud) {
		count := 0
		for key, value := range emp {
			for k, v := range stud {
				if key == k && value == v {
					count++
				}
			}
		}
		if count == len(emp) && count == len(stud) {
			fmt.Println("Both Maps are equal")
		} else {
			fmt.Println("Maps are not equal")
		}
	} else {
		fmt.Println("Maps not equal")
	}
}
