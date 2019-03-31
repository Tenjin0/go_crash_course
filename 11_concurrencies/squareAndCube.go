package main

import (
	"fmt"
	"time"
)

func digits(number int, dchnl chan int) {

	for number != 0 {
		digits := number % 10
		dchnl <- digits
		number /= 10
	}
	close(dchnl)
}

func calcSquares2(number int, squareop chan int) {

	sum := 0
	for number != 0 {
		digits := number % 10
		sum += digits
		number /= 10
	}
	squareop <- sum
}

func calcCubes2(number int, cubeop chan int) {

	sum := 0
	for number != 0 {
		digits := number % 10
		sum += digits
		number /= 10
	}
	cubeop <- sum
}

func calcSquares(number int, squareop chan int) {

	sum := 0
	dch := make(chan int)
	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit
	}

	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {

	sum := 0
	dch := make(chan int)

	go digits(number, dch)

	for digit := range dch {
		sum += digit * digit
	}
	cubeop <- sum
}

func sumSquaresAndCubes(square, cube int) int {

	var start time.Time
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	start = time.Now()

	go calcSquares2(number, sqrch)
	go calcCubes2(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes, time.Now().Sub(start))

	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes = <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes, time.Now().Sub(start))

}
