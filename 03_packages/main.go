package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Tenjin0/go_crash_course/03_packages/geometry/rectangle"
	"github.com/Tenjin0/go_crash_course/03_packages/strutil"
)

var rectLen, rectWidth float64 = 0, 0

func init() {

	var err error
	println("main package initialized")

	for i, val := range os.Args {

		fmt.Printf("Index [%d] = %s\n", i, val)
	}

	if len(os.Args) > 1 {
		rectLen, err = strconv.ParseFloat(os.Args[1], 64)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(os.Args) > 2 {
		rectWidth, err = strconv.ParseFloat(os.Args[2], 64)
		if err != nil {
			log.Fatal(err)
		}

	}

	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {

	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of a rectangle %.2f\n", rectangle.Diagonal(rectLen, rectWidth))

	stringToReverse := "hello"
	fmt.Printf("String to reverse %s: %s\n", stringToReverse, strutil.Reverse(stringToReverse)) // strutilTODO
}
