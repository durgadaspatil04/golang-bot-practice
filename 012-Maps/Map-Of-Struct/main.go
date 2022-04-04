package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	emp1 := employee{
		salary:  1000,
		country: "Scotland",
	}
	emp2 := employee{
		salary:  2000,
		country: "Geramny",
	}
	emp3 := employee{
		salary:  3000,
		country: "India",
	}
	empInfo := map[string]employee{
		"John":  emp1,
		"Bond":  emp2,
		"James": emp3,
	}
	fmt.Println("printing whole map", empInfo)
	//iterate over map
	for name, info := range empInfo {
		fmt.Printf("Employee : %s Salary:%d Country :%s \n", name, info.salary, info.country)
	}
}
