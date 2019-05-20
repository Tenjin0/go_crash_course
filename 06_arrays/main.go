// https://www.thegeekstuff.com/2019/03/golang-slice-examples/
package main

import "fmt"

func changeLocalArray(numbers [3]int) {
	numbers[0] = 0
}

func changeLocalSlice(numbers []int) {
	numbers[0] = 0
	numbers = append(numbers, 1)
}

func changeLocalPointer(numbers *[]int) {
	(*numbers)[0] = 10
	(*numbers) = append((*numbers), 1)
}

func main() {
	numbers := [3]int{1, 2, 3}

	fmt.Println(numbers, len(numbers), cap(numbers))
	changeLocalArray(numbers)
	fmt.Println(numbers, len(numbers), cap(numbers))
	changeLocalSlice(numbers[:])
	fmt.Println(numbers, len(numbers), cap(numbers))
	slice := numbers[:]
	changeLocalPointer(&(slice)) // you cannot append to Array but to slice with pointer
	fmt.Println(numbers, len(numbers), cap(numbers))
	fmt.Println(slice, len(slice), cap(slice))
}
