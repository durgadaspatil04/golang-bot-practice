package main

import "fmt"

func display(chen chan<- int) { //unction converts this channel to a send only channel
	chen <- 20
}
func main() {
	ch := make(chan int)
	go display(ch)
	fmt.Println(<-ch)
}
