package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str2 := "Señor"
	LenSp := utf8.RuneCountInString(str2)
	fmt.Println("Length ", LenSp)         //give Actual length - 5
	fmt.Println("No OF Byte ", len(str2)) //give byte size  - 6

}
