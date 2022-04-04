package main

import "fmt"

func totalPrice(Pr int, no int) int { //function Defination
	TotalPr := Pr * no
	return TotalPr
}
func main() {
	var Price = 499
	no := 23
	Total := totalPrice(Price, no) //Function calling
	fmt.Println("Total price=", Total)

}
