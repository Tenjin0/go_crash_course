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
	var lst = [][]int{{69, 130}, {87, 1310}, {30, 40}}
	var common_pgcd int = lst[0][1]
	for i := 1; i < len(lst); i++ {
		common_pgcd = commondenominators.Pgcd(common_pgcd, lst[i][1])
	}
	var lst2 [][]int

	for i := 0; i < len(lst); i++ {

		converted_fraction := []int{lst[i][0], lst[i][1]}

		for j := 0; j < len(lst); j++ {

			coeffmultiplier := lst[j][1] / common_pgcd
			if i != j {
				converted_fraction[0] = converted_fraction[0] * coeffmultiplier
				converted_fraction[1] = converted_fraction[1] * coeffmultiplier
			}
		}
		lst2 = append(lst2, converted_fraction)
	}
	fmt.Println(lst2)

	for i := 0; i < len(lst2); i++ {
		common_pgcd = commondenominators.Pgcd(lst2[i][0], lst2[i][1])
		fmt.Println(common_pgcd)
	}

	// if common_pgcd != 1 {
	// 	for i := 1; i < len(lst2); i++ {
	// 		lst2
	// 	}
	// }

	fmt.Println(lst)
}
