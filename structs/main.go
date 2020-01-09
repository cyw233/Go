package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	// &variable: give the memory address of the value this variable is pointing at
	// jimPointer := &jim

	// Go allows us to use shortcut, so we can just pass jim
	jim.updateName("jimmy")
	jim.print()

}

// *person: a type description means we're working with a pointer to a type, e.g. jimPointer
// *pointerToPerson: an operator that gives the value this memory address is pointing at
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

// Turn address into value with *address
// Turn value into address with &value
