package main

import "fmt"

func fType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("String type value is %s \n", i.(string))
	case int:
		fmt.Printf("Integer type value is %d \n", i.(int))
	default:
		fmt.Printf("Unknown type \n")

	}
}

func main() {
	var num interface{} = 45
	fType(num)
	var s interface{} = "Hello"
	fType(s)
}
