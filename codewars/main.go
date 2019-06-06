package main

import (
	"fmt"

	"github.com/Tenjin0/go_crash_course/codewars/commondenominators"
)

func main() {

	// a1 := [...]string{"arp", "live", "strong"}
	// a2 := [...]string{"lively", "alive", "harp", "sharp", "armstrong"}
	// fmt.Println(rotateformax.MaxRot(56789))
	// fmt.Println(rotateformax.MaxRot2(56789))
	// results := whicharein.InArray(a1[:], a2[:])
	// fmt.Println(results)

	var lst = [][]int{{3, 1}}
	lst2 := commondenominators.ConvertFracts(lst)
	fmt.Println(lst2)
}
