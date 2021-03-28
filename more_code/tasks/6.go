// package main
//
// import "math"
//
// func main() {
// 	const a int = 7
// 	const b float64 = 5.6
// 	const c = a * b
//
// 	x := 8
// 	const xc int = x
//
// 	const noIPv6 = math.Pow(2, 128)
//
// }

package main

// import "math"

func main() {
	// ERROR -> invalid operation: a * b (mismatched types int and float64)
	// const a int = 7
	const a = 7 // make `a` untyped constant
	const b float64 = 5.6
	const c = a * b

	x := 8
	_ = x
	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const xc int = x  // variables belong to runtime

	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const noIPv6 = math.Pow(2, 128) // functions calls belong to runtime

}
