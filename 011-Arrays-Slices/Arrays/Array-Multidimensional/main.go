package main

import "fmt"

func printArray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Println()
	}
}
func main() {
	a := [3][2]string{ //MultiDimensional array with Declaration and Initialization
		{"A", "a"},
		{"B", "b"},
		{"C", "c"},
	}
	fmt.Println(a)

	var b [3][2]int //MultiDimensional array with Declaration
	fmt.Println(b)
	b[0][0] = 1
	b[0][1] = 2
	b[1][0] = 3
	b[1][1] = 4
	b[2][0] = 5
	b[2][1] = 6
	fmt.Println(b)
	fmt.Println("Array printing using range")
	printArray(a)

}
