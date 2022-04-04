//Higher order func is function which take func as argument and return function as output
/*Higher order functions are functions that operate on other functions, either by taking
them as arguments or by returning them.*/
package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func partialSum(x int) func(int) int {

	return func(y int) int {
		return sum(x, y)
	}
}
func main() {
	partial := partialSum(3)
	fmt.Println(partial(7))

}
