package main

import "fmt"

func main() {
	var s1 []int            //slice Without initialised
	s2 := []int{}           //slice Without initialised
	s3 := make([]int, 5, 5) //Initialised slice
	fmt.Println(s1, s2, s3)
	//s1[0] = 1 in both cases of slice s1, s2 can't assign value without inialisation
	//s2[0] = 1
	s3[0] = 1
	fmt.Println(s1, s2, s3)
	s1 = append(s1, 10, 20, 30, 40)
	s2 = append(s2, 9, 8, 7, 4, 5, 6)
	s3 = append(s3, 2, 3, 4, 5, 6)
	fmt.Println(s1, s2, s3)
}
