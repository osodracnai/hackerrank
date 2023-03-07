package hackerrank

import "sort"

func twoArrays(k int32, A []int32, B []int32) string {
	sort.Slice(A, func(i, j int) bool {
		return A[i] < A[j]
	})
	sort.Slice(B, func(i, j int) bool {
		return B[i] < B[j]
	})
	arev := reverse(A)
	for i, _ := range arev {
		if arev[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"
}

func reverse(arg []int32) []int32 {
	result := make([]int32, len(arg))
	for i, value := range arg {
		result[len(arg)-i-1] = value
	}
	return result
}
