package main

import (
	"fmt"
	"time"
)

func Server1(str chan string) {
	time.Sleep(time.Second * 6)
	str <- "from Server 1"
}
func Server2(str chan string) {
	time.Sleep(time.Second * 3)
	str <- "from Server 2"
}

func main() {
	out1 := make(chan string)
	out2 := make(chan string)
	go Server1(out1)
	go Server2(out2)
	select {
	case s1 := <-out1:
		fmt.Println(s1)
	case s2 := <-out2:
		fmt.Println(s2)
	}
}
