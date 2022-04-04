package main

import "fmt"

func main() {
	a := [...]int{10, 20, 30, 40, 50, 60}
	//one of the way declarting array using "..." compiler takes num element init take as len of array
	fmt.Println("Length of array is ", len(a))
	sum := 0

	for index, value := range a {
		//range returns both INDEX and VALUE
		fmt.Println("Index ", index, "Element", value)
		sum += value
	}
	fmt.Println("Sum of array element is", sum)
}
