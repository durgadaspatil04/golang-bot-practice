//IN FLOAT only two types 1) float32.  2)float64
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a, b := 2.65, 4.789 //its type is Default float64
	var c float32       // size 4 byte
	var d float64       //size 8 byte
	fmt.Printf("Type is %T %T\n", a, b)
	c = float32(a + b) // Need Conversion here
	fmt.Println("The Addition is", c)
	fmt.Printf("Type is %T\n", c)
	d = (a + b)
	fmt.Println("The Addition is", d)
	fmt.Printf("Type is %T\n", d)
	fmt.Printf("Size is %d byte %d byte %d byte\n", unsafe.Sizeof(a), unsafe.Sizeof(b),
		unsafe.Sizeof(c))

}
