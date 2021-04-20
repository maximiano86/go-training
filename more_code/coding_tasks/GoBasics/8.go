package main

import "fmt"

func main() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1. Print each variable using a specific verb for its type
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)
	fmt.Printf("score is %#v\n", score)

	// 2. Print the string value enclosed by double quotes 
	fmt.Printf("z is %q\n", z)

	// 3. Print each variable using the same verb 
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)

	// 4. Print the type of y and score
	fmt.Printf("y type: %T, score type: %T\n", y, score)

}
