package main

import "fmt"

func main() {
	var personSalary map[string]int
	if personSalary == nil {
		personSalary = make(map[string]int)
		fmt.Println("map just created", personSalary)
		personSalary["Tenji"] = 35000
		personSalary["Incogito"] = 100000
		fmt.Println(personSalary, "len:", len(personSalary))
		unknownPerson := "Joe"
		value, ok := personSalary[unknownPerson]
		if ok == true {
			fmt.Println("Salary of", unknownPerson, "is", value)
		} else {
			fmt.Println(unknownPerson, "not found")
		}
		fmt.Println("Loop map:")
		for key, value := range personSalary {
			fmt.Printf("personSalary[%s]: %d\n", key, value)

		}

		fmt.Println("Map lack keys, values equality functions though :(")
	}
}
