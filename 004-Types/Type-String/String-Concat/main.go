//string concation different ways
package main

import "fmt"

func main() {
	s1 := "Hello"
	s2 := "World"
	fmt.Printf("address is %v\n", &s1)
	s1 = "Gothere"
	fmt.Printf("address is %v\n", &s1)
	Str := s1 + " " + s2 //0r s1 + s2  this one way to concat
	fmt.Println(Str)
	StrCon := fmt.Sprintf("%s %s", s1, s2) // string concat using Sprintf func its Return a string
	fmt.Println(StrCon)
	s3 := "Hello"
	fmt.Printf("address is %v\n", &(s3))
	fmt.Println(s3)

}
