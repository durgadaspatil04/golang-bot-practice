//Printing Char of string Using Rune
package main

import "fmt"

func printChar(s string) {
	sl := []rune(s) // string converted Slice or rune
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%c", sl[i])
	}
	fmt.Printf("\n Bytes:")
	for i := 0; i < len(sl); i++ {
		fmt.Printf("%x", sl[i])
	}
}

func main() {
	s := "Señor" /*we Use rune here bcoz string's
	some character takes 2 bytes so WE need to use RUNE(4 byte) not BYTE(1 byte)*/
	printChar(s)

}
