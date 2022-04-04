/*It is perfectly legal for a function to return a pointer of a local variable.
The Go compiler is intelligent enough and it will allocate this variable on the heap.*/
package main

import "fmt"

func hello() *int {
	i := 45
	return &i
}
func main() {
	d := hello()
	fmt.Println("Value of d", *d)

}
