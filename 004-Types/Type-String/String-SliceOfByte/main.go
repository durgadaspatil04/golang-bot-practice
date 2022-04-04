package main

import "fmt"

func main() {
	Sb := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(Sb) //conversion of byte to string
	fmt.Println("String is", str)
	Sb2 := []rune{0x4a, 0x6f, 0x68, 0x6e, 0x20, 0x53, 0x6e, 0x6f, 0x77}
	str2 := string(Sb2) //conversion of rune to string
	fmt.Println("String is", str2)

}
