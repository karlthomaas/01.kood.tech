package piscine

func IsUpper(s string) bool {
	for _, char := range s {
		if char < 65 || char > 90 {
			return false
		}
	}
	return true
}
