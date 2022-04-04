package main

import "fmt"

func main() {
	var new chan int                 //channel declaration using var
	anotherChan := make(chan string) //channel declaration using make function
	fmt.Println(new)
	fmt.Println(anotherChan) //this is initialised channel
	if new == nil {
		fmt.Println("The new channel is nil going to define it")
		new = make(chan int) //initialisation of channel
		fmt.Printf("The type of channel is- %T\n", new)
	}
}
