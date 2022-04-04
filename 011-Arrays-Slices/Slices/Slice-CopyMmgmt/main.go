package main

import "fmt"

func test(sl []int) {
	sl[2] = 11
}
func main() {
	original := []int{1, 2, 3, 4, 5}
	sliced := original[:len(original)-1] //slicing slice
	ncopy := make([]int, len(sliced))    //initialised new slice
	copy(ncopy, sliced)                  //copy the sliced(slice refe) into ncopy(destination)
	fmt.Println("orginal slice-", original)
	fmt.Println("slicing of slice s-", sliced)
	test(ncopy)
	fmt.Println("After func call", sliced) //the changes made into the copy slice not into slicing slice
	fmt.Println("After func call s3", ncopy)
}
