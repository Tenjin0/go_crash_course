package strutil

import (
	"fmt"
)

func Reverse(s string) string {

	runes := []rune(s)
	fmt.Println(runes, len(runes))
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %c, j: %c\n", s[i], s[j])
	}
	return string(runes)
}
