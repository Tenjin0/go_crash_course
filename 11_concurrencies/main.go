package main

import (
	"fmt"
	"time"
)

func numbers() {
	for i := 0; i < 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i < 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {

	// go numbers()
	// go alphabets()
	// fmt.Println("main function: after go call")

	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("main terminated")

	// channel()
	// square, cube := 500, 200
	// fmt.Printf("square: %d + cube: %d = %d\n", square, cube, sumSquaresAndCubes(square, cube))

	// Buffered_channel()
	// WaitGroup()
	// ExecuteWorkerPool(10, 10)
	// Select()
	CritalSectionWithoutMutex()
	CritalSectionWithMutex()
	CriticalSectionWithChan()
}
