package piscine

func IsLower(s string) bool {
	for _, char := range s {
		if char < 97 || char > 122 {
			return false
		}
	}
	return true
}
