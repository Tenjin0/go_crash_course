package main

import (
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

type areaError2 struct {
	err    string
	length float64
	width  float64
}

func (e *areaError2) Error() string {
	return e.err
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNegative() bool {
	return e.width < 0
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		// return 0, errors.New("Area calculation failed, radius is less than 0")
		// return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
		return 0, &areaError{"radius is negative", radius}
	}

	return math.Pi * radius * radius, nil

}

func areaRect(length, width float64) (float64, error) {

	err := ""

	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if len(err) >= 0 {
			err += ", "
		}
		err += "width is less than zero"
	}

	if len(err) > 0 {
		return 0, &areaError2{err, length, width}
	}
	return length * width, nil

}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Println(err)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)

	length, width := -5.0, -9.0
	areaRect, errRect := areaRect(length, width)
	if errRect != nil {
		if errRect, ok := errRect.(*areaError2); ok {
			if errRect.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", errRect.length)
			} else if errRect.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", errRect.width)
			}
		}
		return
	}

	fmt.Printf("area of the retangle:  %0.2f", areaRect)
}
