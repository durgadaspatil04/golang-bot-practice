//single diresction channel with error
package main

func sendData(ch chan<- int) {
	ch <- 20
}
func main() {
	ch := make(chan<- int)
	go sendData(ch)
	//fmt.Println(<-ch)
}
