package main

import "fmt"

func main() {
	var Sid [5]int                   //Declaration of array
	name := [5]string{"a", "b", "c"} //array short Declaration and Initialization
	//IN short Declaration Not compulsory to Assign all values
	age := [3]int{}
	age[0] = 1
	//age[4] = 1 it shows out of bound error
	fmt.Println("age array", age)   //prints whole array
	fmt.Println("name array", name) //prints whole array
	Sid[0] = 1                      //assigning values
	Sid[1] = 2
	Sid[2] = 3
	fmt.Println("Sid array 1st element", Sid[0]) //print single value
	fmt.Println("Sid array", Sid)                //prints whole array
}
