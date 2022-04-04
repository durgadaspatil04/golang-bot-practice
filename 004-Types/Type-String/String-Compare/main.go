package main

import "fmt"

func StringCompare(s1, s2 string) {
	if s1 == s2 {
		fmt.Printf("string %s and %s are equal", s1, s2)
		return
	}
	fmt.Println("Strings are Not equal")
}
func main() {
	str1 := "John"
	str2 := "Snow"
	StringCompare(str1, str2)
	s1 := "hello"
	s2 := "hello"
	StringCompare(s1, s2)
}
