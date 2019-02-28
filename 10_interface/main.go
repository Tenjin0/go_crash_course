package main

import "fmt"

type MyString string

type VowelsFinder interface {
	FindVowels() []rune
}

type SalaryInterface interface {
	setSalary(salary int)
	getSalary() int
}

type Person struct {
	name   string
	salary int
}

func (p *Person) setSalary(salary int) {
	p.salary = salary
}

func (p Person) getSalary() int {
	return p.salary
}

func printSalary(s SalaryInterface) {
	fmt.Println("Salary", s.getSalary())
}

func find(element rune, listElement []rune) bool {

	for i := 0; i < len(listElement); i++ {
		if listElement[i] == element {
			return true
		}
	}
	return false
}

func (ms MyString) FindVowels() []rune {

	var vowels []rune

	for _, character := range ms {

		if character == 'a' || character == 'e' || character == 'o' || character == 'u' || character == 'i' {
			if !find(character, vowels) {
				vowels = append(vowels, character)
			}
		}

	}
	return vowels
}

type I interface {
	M(name string)
}
type T struct {
	name string
}

func (t *T) M(name string) {
	t.name = name
}

func describe(i interface{}) {

	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {

	s, ok := i.(int)
	if ok {
		fmt.Println(s)
	}
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("It is a string: %s\n", i)
	case int:
		fmt.Printf("It is an int: %d\n", i)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {

	name := MyString("coucou")
	var v VowelsFinder
	v = name
	fmt.Printf("%c\n", name.FindVowels())
	fmt.Printf("%c\n", v.FindVowels())
	var patrice SalaryInterface = &Person{
		name: "Patrice",
	}

	f := SalaryInterface.setSalary
	f(patrice, 35000)
	original, ok := patrice.(*Person)
	if ok {
		fmt.Println(original.salary)

	}
	printSalary(patrice)

	var thomas SalaryInterface

	thomas = &Person{
		name: "Thomas",
	}

	thomas.setSalary(30000)
	printSalary(thomas)

	var s interface{} = "Hello World"
	var i interface{} = 55

	describe(s)
	describe(i)
	assert(s)
	assert(i)
	findType(s)
	findType(i)
	findType(89.98)
}
