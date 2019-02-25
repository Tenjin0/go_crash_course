package main

import "fmt"

func test_map() {

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

func bytesToString() {

	byteSlice := []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	str := string(byteSlice)
	fmt.Println(str)

}

func runesToString() {

	runesSlice := []rune{0x48, 0x65, 0x6c, 0x6c, 0x6f}
	str := string(runesSlice)
	fmt.Println(str)

}
func test_strings() {

	hello := "Hello World"

	for i := 0; i < len(hello); i++ {

		fmt.Printf("%x ", hello[i])
	}
	fmt.Printf("\n")

	helloRunes := []rune(hello)

	for _, helloRune := range helloRunes {
		fmt.Printf("%c ", helloRune)
		fmt.Printf("%d ", helloRune)
	}
	fmt.Printf("\n")

	bytesToString()
	runesToString()
	fmt.Printf("Outside func, before change %p %c\n", helloRunes, helloRunes)
	mutateString(helloRunes)
	fmt.Printf("Outside func, after change %p %c\n", helloRunes, helloRunes)

}
func mutateString(s []rune) string {
	fmt.Printf("Inside func, before change %p %c\n", s, s)
	s[0] = 'a' // single quote represents a rune
	s = append(s, '!')
	fmt.Printf("Inside func, after change %p %c\n", s, s)
	return string(s)
}
func main() {

	// test_map()
	test_strings()
}
