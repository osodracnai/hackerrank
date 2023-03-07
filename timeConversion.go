package hackerrank

import "time"

func timeConversion(s string) string {
	t, err := time.Parse("03:04:05PM", s)
	if err != nil {
		panic(err)
	}
	return t.Format("15:04:05")
}
