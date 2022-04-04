package main

import (
	"fmt"
	"time"
)

func display() {
	fmt.Println("Coming to Goroutine Function")
}

func main() {
	go display() //goroutine started
	/*called the "Sleep method" of the "time package" which sleeps the go routine in which
	it is being executed. In this case the main goroutine is put to sleep for 1 second.*/
	time.Sleep(time.Second * 1) //Jya goroutine madhe excute hote tya goroutine la sleep karte
	fmt.Println("This Main Function")
}
