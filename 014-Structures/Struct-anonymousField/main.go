package main

import "fmt"

type student struct {
	int //this are anonymous Fields with no name
	string
}

func main() {
	s1 := student{
		string: "Dsp",
		int:    20,
	}
	fmt.Println("Anonymous field access is-", s1.int, s1.string)
}
