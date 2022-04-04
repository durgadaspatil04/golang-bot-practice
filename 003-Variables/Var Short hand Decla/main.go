//program for short hand declaration use this kind of method to declare a variable as much as possible
package main

import "fmt"

func main() {
	age := 10 //short hand declaration using this u can limit its scope
	fmt.Println("Age=", age)
	for i := 0; i < 2; i++ {
		count := 1
		fmt.Println("count=", count)

	}
	// "count" if u try to access count variable outside the its give error so its scope is limited
}
