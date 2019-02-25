package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

var person struct { // anonymous structures
	firstname, lastname string
	age                 int
}

func main() {
	emp := Person{
		firstname: "Tenji",
		lastname:  "PETIT",
		age:       38,
	}
	// emp2 := struct person{"Toto", "Dupond", 50}

	fmt.Println(emp)
}
