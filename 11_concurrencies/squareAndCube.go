package main

func calcSquares(num int, squareop chan int) {
	sum := 0
	for i := 0; i < num; i++ {
		sum += i * i
	}
	squareop <- sum

}

func calcCubes(num int, cubeop chan int) {

	sum := 0
	for i := 0; i < num; i++ {
		sum += i * i * i
	}

	cubeop <- sum
}

func sumSquaresAndCubes(square, cube int) int {

	var op chan int
	op = make(chan int)

	go calcSquares(square, op)
	go calcCubes(cube, op)

	squares, cubes := <-op, <-op

	return squares + cubes
}
