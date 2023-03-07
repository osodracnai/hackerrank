package hackerrank

func flippingMatrix(matrix [][]int32) int32 {

	for _, a := range matrix {
		b := reverse(a)
		for y, _ := range a {
			print(y)
			print(b)
		}
	}
	return -1
}

//func sum(matrix [][]int32) int32 {
//	sumlenght := len(matrix) / 2
//	fmt.Printf("sumlenght %v \n", sumlenght)
//	for i, a := range matrix {
//		for y, _ := range a {
//
//		}
//	}
//	return -1
//}

//func reverse(arg []int32) []int32 {
//	result := make([]int32, len(arg))
//	for i, value := range arg {
//		result[len(arg)-i-1] = value
//	}
//	return result
//}
