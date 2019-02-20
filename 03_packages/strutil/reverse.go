package strutil

import "fmt"

func Reverse(s string) string {

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		fmt.Printf("i: %b, j: %b\n", s[i], s[j])
	}
	return s
}
