package main

import "fmt"

var age int32 = 37

func main() {

	const isCool = true
	name, email := "Tenji", "patricepetit@gmail.com"
	fmt.Println(name, email)
	fmt.Printf("isCool: %T\n", isCool)
}
