package main

import "fmt"

func main() {
	num := 20
	var pt *int     //pointer Declaration
	fmt.Println(pt) //pointer zero value is Nil
	pt = &num       //pointer initialisation
	fmt.Printf("Type of the pt- %T and Address Stored in pt=%d and Value At PT %d", pt, pt, *pt)
}
