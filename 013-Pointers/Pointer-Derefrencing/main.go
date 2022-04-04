/*Dereferencing a pointer means accessing the value of the variable to which the pointer
points. *a is the syntax to deference a.*/

package main

import "fmt"

func main() {
	a := 35
	b := &a
	fmt.Printf("The Address of 'a'- %d\n", b)
	fmt.Printf("The Value At 'a'- %d\n", *b)
}
