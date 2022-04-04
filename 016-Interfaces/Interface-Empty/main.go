package main

import "fmt"

//here we used empty interface as a function parameter so it take any TYPE value
func info(empty interface{}) {
	fmt.Printf("Value= %v and Type is-%T\n", empty, empty)
}

func main() {
	str := "James Bond"
	info(str)
	num := 34
	info(num)
	strt := struct {
		name string
	}{name: "MoneyPenny"}
	info(strt)

}
