package main

import "fmt"

func main() {
	//var age int = 21
	var age = 21
	fmt.Println("Age:", age)

	dir := "house"
	// blank identifier for declared but not used variables
	_ = dir

	name := "Maximiano"
	fmt.Println("Name:", name)

	car, cost := "VW", 25000
	fmt.Println("Car:", car, "\nCost:", cost)

	var (
		salary    float64
		firstName string
		gender    bool
	)

	_, _, _ = salary, firstName, gender

	sum := 5 + 2.3

	fmt.Println(sum)
}

// will error out can only used short declartion at local variables
/*
package main
import "fmt"
y := 20
func main() {
	x := 10
	fmt.Println(x)
}
*/
