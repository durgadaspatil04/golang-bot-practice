package main

import "fmt"

type salaryCalculator interface {
	calculateSalary() float64
}
type permanent struct {
	empId    int
	basicPay float64
	pf       float64
}

type contract struct {
	empId    int
	basicPay float64
}

func (p permanent) calculateSalary() float64 {
	salary := p.basicPay + p.pf
	return salary
}
func (c contract) calculateSalary() float64 {
	salary := c.basicPay
	return salary
}

func totalSalaryExpense(s []salaryCalculator) {
	expense := float64(0)
	for _, v := range s {
		expense = expense + v.calculateSalary()
	}
	fmt.Printf("Total Expense Per Month %v\n", expense)
}

func main() {
	p1 := permanent{1, 2000, 300}
	p2 := permanent{2, 3400, 340}
	c1 := contract{3, 3000}
	c2 := contract{4, 2500}
	employee := []salaryCalculator{p1, p2, c1, c2}
	totalSalaryExpense(employee)
}
