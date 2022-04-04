package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Nilesh Congrats"
	r := []rune(str)
	fmt.Println("String length", utf8.RuneCountInString(str))
	count := 0
	for i, _ := range r {
		i++
		count++
	}
	fmt.Println("Length OF string", count)
}
