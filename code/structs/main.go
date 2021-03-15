package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)

	max := person{
		firstName: "Max",
		lastName:  "Hernandez",
		contactInfo: contactInfo{
			email:   "max@gmail.com",
			zipCode: 44860,
		},
	}

	// gives the adress in memory
	// maxPointer := &max
	// maxPointer.updateName("Alex")
	// max.print()

	// no need to creat pointer if the receiver is a pointToX
	max.updateName("Alex")
	max.print()

}

// we are working with  pointer to a person type
func (pointerToPerson *person) updateName(newFirstName string) {
	// operator manipulate the value in this pointer memory address
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
