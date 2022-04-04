package main

import "fmt"

func main() {
	employee := make(map[string]int) //declaration and initialisation with nil
	fmt.Println("Map is", employee)  //output is nil-empty map
	//adding values to the map
	employee["john"] = 12000
	employee["james"] = 1000
	employee["bond"] = 1000 //same key value gets overwrite bcoz key should be unique
	employee["bond"] = 20000
	fmt.Println("value added new map is", employee)
}
