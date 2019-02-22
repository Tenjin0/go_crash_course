package main

import (
	"fmt"
)

func calculBill(price, no int) int {

	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (area float64, perimeter float64) {

	area = length * width
	perimeter = (length + width) * 2
	return
}

func main() {

	area, perimeter := rectProps(10.8, 5.6)

	fmt.Printf("total price %d\n", calculBill(10, 5))
	fmt.Printf("Area %f, Perimeter %f\n", area, perimeter)

}
