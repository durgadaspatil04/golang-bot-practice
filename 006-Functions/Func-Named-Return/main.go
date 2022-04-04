package main

import "fmt"

func test(s int) (t int) {
	t = s * 2
	return
}
func main() {
	s := 15
	fmt.Println(test(s))
}
