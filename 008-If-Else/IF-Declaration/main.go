package main

import "fmt"

func main() {
	if r := 10; r == 10 {
		//assigining and condition in if stmt scope of r limited to if block
		fmt.Println(r)
	}
}
