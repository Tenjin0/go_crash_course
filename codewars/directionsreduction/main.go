// https://www.codewars.com/kata/directions-reduction/train/go
package main

import "fmt"

func isTwoDirectionOpposite(first, second string) bool {

	return first == "NORTH" && second == "SOUTH" ||
		first == "SOUTH" && second == "NORTH" ||
		first == "EAST" && second == "WEST" ||
		first == "WEST" && second == "EAST"
}

func removeElements(arr []string, i int) []string {
	return append(arr[:i], arr[i+2:]...)
}

func DirReduc(arr []string) []string {
	// your code
	var i int
	for {
		if i+1 < len(arr) && isTwoDirectionOpposite(arr[i], arr[i+1]) {
			arr = removeElements(arr, i)

			if i > 0 {
				i--
				continue
			}
		}
		if i < len(arr)-1 {
			i++
		} else {
			break
		}
	}
	return arr[:]
}

func main() {
	entryArr := []string{"NORTH", "SOUTH", "SOUTH", "EAST", "WEST", "NORTH", "WEST"}

	outArr := DirReduc(entryArr)

	fmt.Println(outArr)
}
