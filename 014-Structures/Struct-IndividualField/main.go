package main

import "fmt"

type stud struct {
	name    string
	age, id int
}

func main() {
	stud1 := stud{"James", 23, 01}
	//accessing individual fields of structure using DOT(.) Operator

	fmt.Println("Student ID-", stud1.id)
	fmt.Println("Name-", stud1.name)
	fmt.Println("Age-", stud1.age)
	stud1.name = "MoneyPenny" //assigning element A value

	fmt.Println("Individual fileds of struct student-", stud1.id, stud1.name, stud1.age)
}
