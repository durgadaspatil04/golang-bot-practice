/*The new function takes a type as an argument and returns a pointer to a newly
allocated zero value of the type passed as argument.*/
package main

import "fmt"

func main() {
	ptr := new(int) //Pointer declaration using new func
	fmt.Printf("ptr value is %d, type is %T, address is %v\n", *ptr, ptr, ptr)
	*ptr = 45
	fmt.Printf("The new Value of PTR %d\n ", *ptr)

}
