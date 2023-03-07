package hackerrank

import (
	"fmt"
	"sort"
)

func lonelyinteger(a []int32) int32 {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	if len(a) == 1 {
		return a[0]
	}
	var result int32

	for i := 1; i < len(a); i++ {
		if a[i-1] == a[i] {
			continue
		}
		if i >= len(a)-1 {
			result = a[i]
			break
		}
		if a[i] == a[i+1] {
			continue
		}

		result = a[i]
	}
	return result
}
func lonelyinteger2(a []int32) int32 {
	resp := map[int32]int32{}
	for _, value := range a {
		resp[value] += 1
	}
	for key, value := range resp {
		fmt.Print(key)
		if value == 1 {
			return key
		}
	}
	return -1
}
