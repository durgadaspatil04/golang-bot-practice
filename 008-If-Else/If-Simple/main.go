package main

import "fmt"

func main() {
	num := 20
	if num%2 == 0 {
		fmt.Println("The Number", num, "is Even")
		return
	}
	fmt.Println("The Number is odd")

}
