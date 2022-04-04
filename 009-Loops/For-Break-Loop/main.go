package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i > 4 {
			break // breaks the normal loop excution and exit the loop
		}
	}
}
