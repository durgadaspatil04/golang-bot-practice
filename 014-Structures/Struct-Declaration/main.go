package main

import "fmt"

type employee struct { //struct declaration
	fname string
	last  string
	age   uint8
}

func main() {
	type EmpNew struct { //struct declaration
		fname, last string
		age         uint8
	}
	//creating struct specifying field names
	eOld := employee{
		fname: "Dsp",
		last:  "patil",
		age:   20,
	}
	//creating struct without specifying field names
	eNew := EmpNew{"James", "Bond", 54}

	fmt.Printf("Old Struct %v\n New Struct is %v\n", eOld, eNew)
	fmt.Println("Old Struct is- ", eOld.fname, eOld.last, eOld.age)

}
