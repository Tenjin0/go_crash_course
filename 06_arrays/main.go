package main

import (
	"fmt"
	"math/rand"
	"time"
)

func changeLocal(num [3]int) {
	num[0] = 4
	// num = append(num, 3)

}
func changeLocalSlice(num []int) {
	num[0] = 4
	num = append(num, 3) // will not work -> copy len and cap and a pointer to the array

}
func changeLocalWithPointer(num *[]int) {
	(*num)[0] = 90
	(*num) = append(*num, 3)
}

func appendToSliceArray(num []int) {
	fmt.Printf("inside function before append g %v, length: %d ,capacity: %d, &: %p\n", num, len(num), cap(num), &num)
	num = append(num, 1)
	num[0] = 99
	fmt.Printf("inside function after append g %v, length: %d ,capacity: %d, &: %p\n", num, len(num), cap(num), &num)

}

func main() {

	var a [3]int
	b := [...]int{1, 2, 2}
	c := b
	d := b[:]

	b[2] = 3

	for i := 0; i < len(a); i++ {
		rand.Seed(time.Now().UnixNano())
		a[i] = rand.Intn(10 + 1)
	}

	e := make([]int, 5, 5)

	e = append(e, 1) // new array is created with double capacity

	f := append(d, d...)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("b:", b[:], "slice of b")
	changeLocal(b)
	fmt.Println("b:", b, "change inside function")
	changeLocalSlice(b[:])
	fmt.Println("b:", b, "change inside function (slice)")

	fmt.Println("c:", c, "assigned by value")
	fmt.Println("d:", d, "assigned by reference")
	changeLocalWithPointer(&d)
	fmt.Println("d:", d, "change inside function (pointer)")
	fmt.Println("e:", e, "created by make", "length", len(e), " ,capicity:", cap(e))
	e = e[:len(e)-1]
	fmt.Println("e:", e, "remove last item", "length", len(e), " ,capicity:", cap(e))
	fmt.Println("f", f, "d + d", "work only with slice")

	for i, v := range a { //ignores index
		fmt.Printf("a[%d]: %d", i, v)
	}
	fmt.Println()

	g := make([]int, 4, 5)
	fmt.Printf("before function before append g %v, length: %d ,capacity: %d, &: %p\n", g, len(g), cap(g), &g)
	appendToSliceArray(g)
	fmt.Printf("after function before append g %v, length: %d ,capacity: %d, &: %p\n", g, len(g), cap(g), &g)
}
