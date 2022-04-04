package main

import "fmt"

func main() {
	first := []string{"a", "b", "c"}
	second := []string{"X", "Y", "Z"}
	Alpha := append(first, second...) //Appending One slice to Another but Use veridiac...
	fmt.Println("Appended Slice", Alpha)
}
