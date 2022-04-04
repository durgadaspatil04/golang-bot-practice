package main

import (
	"Pack-Simple-App/SimpleInt"
	"fmt"
)

func main() {
	fmt.Println("The Simple interest calculation")
	var p, r, t float64 = 12, 2, 22
	int := SimpleInt.SimpleInterest(p, r, t)
	fmt.Println("Interest is", int)
}
