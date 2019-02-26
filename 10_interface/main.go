package main

import "fmt"

type MyString string

type VowelsFinder interface {
	FindVowels() []rune
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

func main() {

	name := MyString("coucou")
	var v VowelsFinder
	v = name
	fmt.Printf("%c\n", name.FindVowels())
	fmt.Printf("%c\n", v.FindVowels())
}
