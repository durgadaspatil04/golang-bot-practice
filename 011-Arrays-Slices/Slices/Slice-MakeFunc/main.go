package main

import "fmt"

//sl here initialised slice and sl2 not inialised
//Make function initialised with zero values
func main() {
	sl := make([]int, 4, 6) //make([]T, len, cap) declaration and initialization
	var sl2 []int
	sl[0] = 1
	fmt.Println(sl)
	//sl2[0] = append(sl2, 2) the slice here not initialised so you cannot assign the value to it
	fmt.Println(sl2)
}
