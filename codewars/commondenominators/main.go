package commondenominators

import (
	"fmt"
	"strconv"
)

// kata: https://www.codewars.com/kata/54d7660d2daf68c619000d95
// example: https://www.youtube.com/watch?v=WdsvijS5MRA
// lcm: https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/

// Pgcd is used to find the pgcd between two numbers
func GCD(a, b int) int {

	var sup, inf int

	if a > b {
		sup = a
		inf = b
	} else if b > a {
		sup = b
		inf = a
	} else {
		return a
	}

	for inf != 0 {
		t := inf
		inf = sup % inf
		sup = t
	}

	return sup
}

func LCM(a, b int, integers ...int) int {

	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}
	return result
}

func SameDenominator(lst [][]int) [][]int {

	var denominators []int

	for i := 0; i < len(lst); i++ {
		gcd := GCD(lst[i][0], lst[i][1])
		fmt.Println(gcd, lst[i])
		lst[i][0] = lst[i][0] / gcd
		lst[i][1] = lst[i][1] / gcd
		denominators = append(denominators, lst[i][1])
	}

	if len(lst) == 1 {
		return lst
	}
	a := denominators[0]
	b := denominators[1]
	lcm := LCM(a, b, denominators[2:]...)

	var lst2 [][]int

	var result string = ""
	for _, fraction := range lst2 {
		result += "(" + strconv.Itoa(fraction[0]) + "," + strconv.Itoa(fraction[1]) + ")"
	}
	for i := 0; i < len(lst); i++ {

		convertedFraction := []int{lst[i][0], lst[i][1]}
		coeffmultiplier := lcm / convertedFraction[1]
		convertedFraction[0] = convertedFraction[0] * coeffmultiplier
		convertedFraction[1] = convertedFraction[1] * coeffmultiplier

		lst2 = append(lst2, convertedFraction)
	}
	return lst2
}

func ConvertFracts(a [][]int) string {
	// your code
	lst2 := SameDenominator(a)
	var result string = ""
	for _, fraction := range lst2 {
		result += "(" + strconv.Itoa(fraction[0]) + "," + strconv.Itoa(fraction[1]) + ")"
	}
	return result
}
