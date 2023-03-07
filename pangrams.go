package hackerrank

import "strings"

func pangrams(s string) string {
	lower := strings.ToLower(s)
	base := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, char := range base {
		if !strings.Contains(lower, char) {
			return "not pangram"
		}
	}
	return "pangram"
}
