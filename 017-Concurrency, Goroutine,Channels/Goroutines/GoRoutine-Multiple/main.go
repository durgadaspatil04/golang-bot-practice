package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 1; i < 6; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("%d\n", i)
	}
}

func alphabets() {
	for j := 'a'; j < 'h'; j++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c\n", j)
	}
}

func main() {
	//here we started multiple goroutine
	go numbers()
	go alphabets()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("Main Executed")
}
