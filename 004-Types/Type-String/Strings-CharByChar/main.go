package main

import (
	"fmt"
)

func printChar(s string) {
	fmt.Printf("String is %s\n", s)
	fmt.Printf("Bytes-")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")
	fmt.Printf("String Charcters:-")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func main() {
	name := "Rajveer Patil"
	printChar(name)
}
