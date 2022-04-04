package main

import "fmt"

func main() {
	s1 := new([11]int)[0:10] //Declaration and initialisation of slice init with zero value
	fmt.Println(s1)
	s1[0] = 1
	s1[1] = 11
	s1[2] = 22
	fmt.Println(s1)
}
