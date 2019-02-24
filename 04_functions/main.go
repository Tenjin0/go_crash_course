package main

import (
	"fmt"
	"math/rand"
	"time"
)

func calculBill(price, no int) int {

	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (area float64, perimeter float64, foo int) {

	area = length * width
	perimeter = (length + width) * 2
	foo = 1
	return
}

func random(absMax int) int {

	rand.Seed(time.Now().UnixNano())
	var randNumber = rand.Intn(absMax+1) - absMax/2
	return randNumber

}

func main() {

	area, perimeter, _ := rectProps(10.8, 5.6)

	fmt.Printf("total price %d\n", calculBill(10, 5))
	fmt.Printf("Area %f, Perimeter %f\n", area, perimeter)
	fmt.Printf("random: %d\n", random(2))
}
