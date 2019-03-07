package main

import "fmt"

func largest(nums []int) {

	max := nums[0]

	finished := func() {
		fmt.Println("Finished: Largest number in", nums, "is", max)
	}

	defer finished()

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

}

func deferStack() {

	name := "Maxime"
	fmt.Println("Original name:", name)
	fmt.Print("Defer stack of print characters: ")
	defer fmt.Println()
	for _, c := range []rune(name) {
		defer fmt.Printf("%c", c)
	}

}

func main() {
	largest([]int{78, 109, 2, 563, 300})
	deferStack()
}
