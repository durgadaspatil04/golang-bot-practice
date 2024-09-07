package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// as uint8 range is 255 so u cannot store value above 255 error Overflow occurs
	var sizeCm uint8 = 255 // cannot use 256 (untyped int constant) as uint8 value in variable declaration (overflows)
	var a int
	fmt.Println("Value is", sizeCm)
	fmt.Printf("Type is %T\n", sizeCm)                           //prints data type
	fmt.Printf("The Size of var is %d\n", unsafe.Sizeof(sizeCm)) //prints the size - 1byte
	fmt.Printf("Type is %T\n", a)                                //prints data type
	fmt.Printf("The Size of var is %d\n", unsafe.Sizeof(a))      //prints the size- 8byte
}
