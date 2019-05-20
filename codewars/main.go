package main

import (
	"fmt"

	"github.com/Tenjin0/go_crash_course/codewars/aretheythesame"
)

func main() {

	// a1 := [...]string{"arp", "live", "strong"}
	// a2 := [...]string{"lively", "alive", "harp", "sharp", "armstrong"}
	// fmt.Println(rotateformax.MaxRot(56789))
	// fmt.Println(rotateformax.MaxRot2(56789))
	// results := whicharein.InArray(a1[:], a2[:])
	// fmt.Println(results)
	var a1 = []int{}
	var a2 = []int{25, 49}
	fmt.Println(aretheythesame.Comp(a1, a2))
}
