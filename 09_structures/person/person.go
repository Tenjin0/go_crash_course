package person

import "fmt"

type Person struct {
	Firstname string
	Lastname  string
	Age       int
	Address   Address
}

type Person2 struct {
	Firstname string
	Lastname  string
	Age       int
	Address
}

var person struct { // anonymous structures
	firstname, lastname string
	age                 int
}

type Address struct {
	City, Country string
}

func (a Address) DisplayAddress() {

	fmt.Println("Address", a)
}

func (p *Person) ChangeAge(age int) {
	p.Age = age
}
