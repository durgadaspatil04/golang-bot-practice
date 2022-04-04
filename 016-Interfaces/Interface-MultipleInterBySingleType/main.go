package main

import "fmt"

type bothInter interface {
	EmployeeInfo
	leaveCalculator
}
type EmployeeInfo interface {
	salaryEmp() int
}

type leaveCalculator interface {
	leave() int
}

type employee struct { //this type implementing multiple interfaces
	name, Lname           string
	empId, salary, pf, lv int
}

func (e employee) salaryEmp() int {
	sal := e.salary + e.pf
	return sal
}

func (e employee) leave() (l int) {
	l = e.lv
	return
}

func main() {
	var i EmployeeInfo
	e1 := employee{"James", "Bond", 01, 23000, 230, 25}
	i = e1
	income := i.salaryEmp()
	fmt.Println("Salary total", income)
	var i2 leaveCalculator
	i2 = e1
	L := i2.leave()
	fmt.Println("Leaves ", L)
}
