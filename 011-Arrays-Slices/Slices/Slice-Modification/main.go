package main

import "fmt"

func main() {
	Sl := [...]int{10, 20, 1, 2, 3, 34, 4}
	dSl := Sl[2:5]
	fmt.Println("array before mod of slice", Sl)
	for i := range dSl {
		dSl[i]++
	}
	fmt.Println("Array after mod of slice", Sl) //slice are refernce type the changes
	//made in slice that reflect into the Array
}
