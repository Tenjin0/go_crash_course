package whicharein

import (
	"sort"
	"strings"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func InArray(array1 []string, array2 []string) []string {

	results := make([]string, 0)

	for i := 0; i < len(array1); i++ {
		elementToSearch := array1[i]
		for j := 0; j < len(array2); j++ {

			elementToExamineIn := array2[j]
			if strings.Contains(elementToExamineIn, elementToSearch) && !stringInSlice(elementToSearch, results) {
				results = append(results, elementToSearch)
				break
			}
		}
	}
	sort.Strings(results)
	return results
}
