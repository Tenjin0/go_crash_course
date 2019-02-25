package main

import "fmt"

func change(val *int) {
	*val++
}
func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b", a)
	var c *int
	if c == nil {
		fmt.Println("address of not initialized c", c)
		c = &b
		fmt.Println("b after initialization is", c, *c)
	}
	change(a)
	fmt.Println("b have a new value", b)
}
