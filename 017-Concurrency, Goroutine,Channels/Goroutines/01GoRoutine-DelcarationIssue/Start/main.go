package main

import "fmt"

func display() {
	fmt.Println("Hello To Concurrency Programing")
}
func main() { //The main function runs in its own Goroutine and it's called the main Goroutine.
	go display() //starting goroutine
	//Now the hello() function will run concurrently along with the main() function.
	fmt.Println("I am Main func")
}

/*here func called but didn't get output beacause in go main function is also a goroutine
default so its excution done and program gets exit main and Display() func run concurrently
but because main got excuted first its exit*/
