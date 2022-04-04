//Anonymous function accessing the variable defined outside body.
package main

import "fmt"

func main() {
	l, w := 12, 5

	func() {
		area := l * w //accessing variable declared outside body
		fmt.Println("area", area)
	}()
}
