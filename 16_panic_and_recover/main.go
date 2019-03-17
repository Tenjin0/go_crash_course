package main

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
}

func (p *Person) changeName(firstname *string, lastname *string) {

	defer fmt.Println("deferred call in fullName")
	if firstname == nil {
		panic("runtime error: first name cannot be nil")
	} else {
		p.Firstname = *firstname
	}
	if lastname == nil {
		panic("runtime error: last name cannot be nil")
	} else {
		p.Lastname = *lastname
	}

	fmt.Println("returned normally from fullName")
}
func main() {
	defer fmt.Println("deferred call in main")
	toto := Person{"toto", "Dupont"}
	// newFirstname := nil
	newLastname := "Dupond"
	toto.changeName(nil, &newLastname)
	fmt.Println(toto)
}
