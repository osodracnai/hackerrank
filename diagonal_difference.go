package hackerrank

import "math"

func diagonalDifference(arr [][]int32) int32 {
	var leftToRight int32 = 0
	var rightToLeft int32 = 0

	for x, a := range arr {
		for y, _ := range a {
			if x == y {
				leftToRight += arr[x][y]
			}
			if x+y == len(a)-1 {
				rightToLeft += arr[x][y]
			}
		}
	}
	return int32(math.Abs(float64(leftToRight - rightToLeft)))

}
