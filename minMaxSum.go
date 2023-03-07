package hackerrank

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int64) {
	var max int64
	var min int64
	var sum int64
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	min = arr[0]
	max = arr[0]
	for _, i := range arr {
		sum += i
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	fmt.Printf("%v %v \n", sum-max, sum-min)
}
