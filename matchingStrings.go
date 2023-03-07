package hackerrank

import "strings"

func matchingStrings(stringsV []string, queries []string) []int32 {
	result := make([]int32, len(queries))
	for i, query := range queries {
		for _, str := range stringsV {
			if strings.EqualFold(str, query) {
				result[i] += 1
			}
		}
	}
	return result
}
