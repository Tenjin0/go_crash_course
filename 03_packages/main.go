package main

import (
	"fmt"
	"math"

	"github.com/Tenjin0/go_crash_course/03_packages/geometry/rectangle"

	"github.com/Tenjin0/go_crash_course/03_packages/strutil"
)

func main() {

	var rectLen, rectWidth float64 = 6, 7
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))

	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of a rectangle %.2f", rectangle.Diagonal(rectLen, rectWidth))

	strutil.Reverse("hello") // TODO
}
