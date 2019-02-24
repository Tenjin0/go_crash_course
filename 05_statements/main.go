package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func ifelse() {
	if num := rand.Int(); num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
}

func loop() {
	i := 0
	for true {
		if i >= 10 {
			break
		}
		if i%2 == 0 {
			i++
			continue
		}
		fmt.Printf(" %d", i)
		i++

	}
	fmt.Println()
}

func finger(finger int) {
	fmt.Printf("Finger %d is the ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}

}
func letter(letter string) {
	fmt.Printf("letter to check %s: ", letter)
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}
}

func generateLetter() string {

	a := int([]rune("A")[0])
	b := int([]rune("z")[0])
	letter := a + rand.Intn(b-a+1)
	return string(letter)
}

func check1to100(num int) {
	fmt.Printf("number to check %d:\n", num)
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("It is between 0 and 50")
	case num >= 101:
		fmt.Println("It is greeter than 100")
		fallthrough // do not evaluated next case expression though
	case num >= 51:
		fmt.Println("It is greeter than 50")
	default:
		fmt.Println("It is not positive")
	}
}

func main() {

	ifelse()
	loop()
	finger(rand.Intn(6) + 1)
	letter(generateLetter())
	check1to100(rand.Intn(200 + 1))
}
