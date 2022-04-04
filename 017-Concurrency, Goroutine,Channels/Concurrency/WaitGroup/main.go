package main

import (
	"fmt"
	"sync"
	"time"
)

func process(n int, wg *sync.WaitGroup) {
	fmt.Println("The goroutine process started", n)
	time.Sleep(time.Second * 2)
	fmt.Printf("Goroutine %d ended\n", n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	num := 3
	for i := 0; i < num; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
