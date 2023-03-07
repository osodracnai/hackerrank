package hackerrank

func countingSort(arr []int32) []int32 {
	result := make([]int32, 100)
	for _, value := range arr {
		result[value] += 1
	}
	return result
}
