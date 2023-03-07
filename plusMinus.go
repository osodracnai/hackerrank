package hackerrank

import "fmt"

func plusMinus(arr []int32) {
	var positives float32 = 0
	var negatives float32 = 0
	var zeros float32 = 0
	for _, i := range arr {
		if i == 0 {
			zeros += 1
		} else if i > 0 {
			positives += 1
		} else if i < 0 {
			negatives += 1
		}
	}
	length := float32(len(arr))
	fmt.Printf("%.6f \n", positives/length)
	fmt.Printf("%.6f \n", negatives/length)
	fmt.Printf("%.6f \n", zeros/length)
}
