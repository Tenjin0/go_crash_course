// https://www.codewars.com/kata/550498447451fbbd7600041c
package aretheythesame

import "fmt"

func FindIndex(nums []int, search int) int {

	for i, num := range nums {
		if num == search {
			return i
		}
	}
	return -1
}

func Find(nums []int, search int) bool {
	return FindIndex(nums, search) >= 0
}

func Comp(array1 []int, array2 []int) bool {
	fmt.Println(array1)
	fmt.Println(array2)

	if array1 == nil || array2 == nil {
		return false
	}

	if len(array1) != len(array2) {
		return false
	}

	if len(array1) == 0 || len(array2) == 0 {
		return true
	}

	for _, arr1 := range array1 {
		index := FindIndex(array2, arr1*arr1)
		if index < 0 {
			return false
		}
		array2 = append(array2[:index], array2[index+1:]...)
	}

	return len(array2) == 0
}
