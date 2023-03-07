package hackerrank

func birthday(s []int32, d int32, m int32) int32 {
	var result int32 = 0
	for i, _ := range s {
		if i+int(m) > len(s) {
			break
		}
		sumQ := sum(s[i:(i + int(m))])
		if sumQ == d {
			result++
		}

	}
	return result
}
func shift(arg []int32, value int32) []int32 {
	for i, _ := range arg {
		if i == len(arg)-1 {
			arg[i] = value
			break
		}
		arg[i] = arg[i+1]
	}
	return arg
}
func sum(s []int32) int32 {
	var result int32 = 0
	for _, value := range s {
		result += value
	}
	return result
}
