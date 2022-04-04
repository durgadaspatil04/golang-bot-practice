package main

import "fmt"

func main() {
	ch := make(chan string, 2) //buffered channel
	/*Since the channel has a capacity of 2,
	it is possible to write 2 strings into the channel without being blocked.*/
	ch <- "naveen"
	ch <- "paul"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
