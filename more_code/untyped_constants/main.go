package main

import "fmt"

func main() {
	const a float64 = 5.1 // type constant
	const b = 6.4         // untyped const

	const c float64 = a * b
	const str = "Hello" + "Go"

	const d = 5 > 10
	fmt.Println(d)

	// const x int = 5
	// const y float64 = 2.2 * x // strong type, can not multiple diff types

	const x = 5
	const y = 2.2 * x

	var i int = x     // typeless chanegs to int
	var j float64 = x // var j float64 = float64(x)
	_, _ = i, j
}
