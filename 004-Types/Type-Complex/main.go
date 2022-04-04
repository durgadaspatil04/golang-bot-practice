// Complex type has two Types Complex64(4 byte) Complex128(8 byte)
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	c1 := complex(4, 7) // Default Complex type is Complex128
	c2 := 5 + 22i
	var c3 complex64
	add := (c1 + c2)
	multi := c1 * c2
	fmt.Println("Addition is", add)
	fmt.Println("Multiplication is", multi)
	fmt.Printf("C1 type %T\n", c1)
	fmt.Printf("C3 type %T\n", c3)
	fmt.Printf("type %dbyte %d byte %d byte ", unsafe.Sizeof(c1), unsafe.Sizeof(c2), unsafe.Sizeof(c3))

}
