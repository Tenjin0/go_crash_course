package rotateformax

// https://www.codewars.com/kata/56a4872cbb65f3a610000026/solutions/go

import (
	"math"
)

func rotate(digits []int) []int {
	result := make([]int, 0)
	for i := 1; i < len(digits); i++ {
		result = append(result, digits[i])
	}
	result = append(result, digits[0])
	return result
}

func digitsToNumer(digits []int) int64 {

	var result int64

	for i := 0; i < len(digits); i++ {
		result += int64(int(math.Pow10(i)) * digits[len(digits)-i-1])
	}
	return result
}

func concatenateIntWithDigits(digit int, digits []int) []int {
	return append([]int{digit}, digits...)
}

func Max(x, y int64) int64 {
	if x < y {
		return y
	}
	return x
}

func numberToDigits(number int64) []int {
	var digits []int
	digits = make([]int, 0)
	for number != 0 {
		var digit int
		digit = int(number % 10)
		digits = concatenateIntWithDigits(digit, digits)
		number /= 10
	}
	return digits
}

func MaxRot(number int64) int64 {

	max := number

	digits := numberToDigits(number)

	for i := 0; i < len(digits)-1; i++ {
		digits = append(digits[:i], rotate(digits[i:])...)
		max = Max(max, digitsToNumer(digits))
	}

	return max
}
