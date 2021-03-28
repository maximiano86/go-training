package main

import "fmt"

func main() {
	const (
		c1 = iota
		c2 //= iota
		c3 //= iota
	)

	fmt.Println(c1, c2, c3)

	const (
		a = iota * 2
		b
		c
	)

	fmt.Println(a, b, c)

	const (
		x = -(iota + 2)
		_
		y
		z
	)

	fmt.Println(x, y, z)
}
