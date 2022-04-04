package main

import "fmt"

func SquareCal(num int, sqSum chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit
		num /= 10
	}
	fmt.Println(sum)
	sqSum <- sum
}

func CubeCal(num int, cbSum chan int) {
	sum := 0
	for num != 0 {
		digit := num % 10
		sum += digit * digit * digit
		num /= 10
	}
	fmt.Println(sum)
	cbSum <- sum
}

func main() {
	number := 589
	cb := make(chan int)
	sq := make(chan int)

	go SquareCal(number, sq)
	go CubeCal(number, cb)
	square, cube := <-sq, <-cb
	fmt.Println("Final Sum is", square+cube)
}
