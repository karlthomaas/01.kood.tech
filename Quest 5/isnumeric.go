package piscine

func IsNumeric(s string) bool {
	for _, char := range s {
		if string(char) < "0" || string(char) > "9" {
			return false
		}
	}
	return true
}
