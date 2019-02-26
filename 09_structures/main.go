package main

import (
	"fmt"

	"github.com/Tenjin0/go_crash_course/09_structures/person"
)

func main() {

	emp := person.Person{
		Firstname: "Tenji",
		Lastname:  "PETIT",
		Address: person.Address{
			City:    "Paris",
			Country: "France",
		},
	}
	fmt.Println(&emp)
	emp2 := person.Person{"Toto", "Dupond", 50, person.Address{"Paris", "France"}}

	emp3 := &emp2
	(*emp3).Firstname = "Titi"
	fmt.Println(emp2)

	emp4 := person.Person2{"Toto", "Dupond", 50, person.Address{"France", "Paris"}}

	fmt.Println("Direct access to anonymous struct -> promoted field: city", emp4.City)

	emp5 := person.Person{
		Firstname: "Tenji",
		Lastname:  "PETIT",
		Address: person.Address{
			City:    "Paris",
			Country: "France",
		},
	}

	fmt.Println("Stuct are comparable -> emp == emp5:", emp == emp5)

	(&emp).ChangeAge(38)
	fmt.Println(emp.Firstname, "update age:", emp.Age)
	emp.ChangeAge(15)
	fmt.Println(emp.Firstname, "update age:", emp.Age)

	emp.Address.DisplayAddress()

	emp4.DisplayAddress()
}
