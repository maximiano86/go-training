package main

import "fmt"

func main() {
	x, y := 4, 5.1

	z := x * y
	fmt.Println(z)

	a := 5
	b := 6.2 * a
	fmt.Println(b)

}



// AFTER

package main

import "fmt"

func main() {
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5.
	b := 6.2 * a
	fmt.Println(b)

}
