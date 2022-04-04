// String type are Collection of that is Slice of Byte(uint8) in Go
package main

import "fmt"

func main() {
	First := "John"
	Last := "Snow"
	name := First + " " + Last //it concating the string
	Nname := First + Last      //Anothe Way of concating the string
	fmt.Println("The Name is", name)
	fmt.Println("The Name is", Nname)

}
