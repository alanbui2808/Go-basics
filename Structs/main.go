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
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

// this wouldn't work since Go passes by value by default
// p will be copied into the argument instead.
func (p person) updateNameNaive(newFirstName string) {
	p.firstName = newFirstName
}

// *person = pointer of type person
// &p = actual address of p person
// *p = actual value at address p
func (p *person) updateName(newFirstName string) {
	// receiver can be used on p person despite requires *person
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
