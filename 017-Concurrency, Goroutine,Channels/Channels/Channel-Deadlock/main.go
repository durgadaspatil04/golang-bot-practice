package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 23 //write hotay but control flow blocks find for Reader but got nothing so deadlock
	go display(ch)

}
func display(ch chan int) {
	fmt.Println(<-ch)

}
