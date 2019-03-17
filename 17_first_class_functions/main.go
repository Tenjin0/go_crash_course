package main

import (
	"fmt"
)

type op func(a int, b int) int

type student struct {
	firstname string
	lastname  string
	grade     string
	country   string
}

func filter(s []student, f func(student) bool) []student {

	var r []student
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func imap(s []int, f func(int) int) []int {

	var r []int
	for _, v := range s {
		r = append(s, f(v))
	}
	return r
}

func anonymousFunction() {

	a := func() {
		fmt.Println("i'm an anonymous function")
	}

	a()

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

}

func mapExample() {

	a := []int{5, 6, 7, 8, 9}

	b := imap(a, func(n int) int {
		return n + 10
	})
	fmt.Println(b)
}

func simple(a op) {
	fmt.Println(a(60, 7))
}

func appendStr() func(string) string {

	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func studentExample() {

	s1 := student{"toto", "Dupond", "A", "France"}
	s2 := student{"titi", "Dupont", "B", "Russe"}

	s := []student{s1, s2}

	f := filter(s, func(s student) bool {
		if s.grade == "A" {
			return true
		}
		return false
	})

	fmt.Println(f)
}

func main() {

	anonymousFunction()

	var a op = func(a int, b int) int {
		return a + b
	}
	var b op = func(a int, b int) int {
		return a * b
	}

	s := a(5, 6)
	fmt.Println("Sum", s)
	simple(a)
	simple(b)

	c := appendStr()
	d := appendStr()

	fmt.Println(c("World"))
	fmt.Println(d("Everyone"))

	fmt.Println(c("Gopher"))
	fmt.Println(d("!"))

	studentExample()
	mapExample()

}
