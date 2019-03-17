package main

import (
	"fmt"
	"math/rand"
	"reflect"
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
func variadicFun(first int, num ...int) {
	fmt.Printf("first element: %d, rest: %v, type of rest: %T\n", first, num, num)
	num[0] = 99
	num = append(num, 100) // work if not change the length
	num = append(num, 101)
}

func variadicFunWithAnyParams(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}

func main() {

	area, perimeter, _ := rectProps(10.8, 5.6)

	fmt.Printf("total price %d\n", calculBill(10, 5))
	fmt.Printf("Area %f, Perimeter %f\n", area, perimeter)
	fmt.Printf("random: %d\n", random(2))
	a := make([]int, 3, 10)
	b := make([]int, 3, 10)
	variadicFun(1, 2, 3)
	variadicFun(1, a[0:2]...)
	fmt.Println("variadic function with slice passed in parameters", a, "( change slice inside the function!)")
	variadicFun(1, b...)
	fmt.Println("variadic function with slice passed in parameters", b, "( change slice inside the function!)")
	variadicFunWithAnyParams(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}
