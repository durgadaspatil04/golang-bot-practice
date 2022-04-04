package main

import "fmt"

func display(done chan int) {
	fmt.Println("Display Function called demo channel communication")
	done <- 32 //writing to the channel SENDING DATA TO CHANNEL

}

func main() {
	done := make(chan int)
	go display(done)
	/*when data is read from a channel, the read is blocked until some Goroutine
	writes data to that channel.*/
	<-done //reading from the channel or RECEVING DATA FROM CHANNEL
	//receives data from the done channel but does not use or store that data in
	//any variable. This is perfectly legal.
	fmt.Println("this main Executing")
}

/*In line no. 16 (<-done) we are receiving data from the done channel. This line of code
is blocking which means that until some Goroutine writes data to the done channel,
the control will not move to the next line of code. Hence this eliminates the need
for the time.Sleep*/
