package main

import "fmt"

func main() {
	a := [3]string{"Australia", "India", "Italy"}
	b := a //aasign a copy to 'b' so what are changes made in b not refelct in a
	b[0] = "Iraq"
	fmt.Println("Array a", a)
	fmt.Println("Array b", b)
}
