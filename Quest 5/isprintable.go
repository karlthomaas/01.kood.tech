package piscine

func IsPrintable(s string) bool {
	for _, char := range s {
		if string(char) < " " {
			return false
		}
	}
	return true
}
