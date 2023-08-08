package piscine

func BasicAtoi2(s string) int {
	srune := []rune(s)
	result := 0
	// integer unicode stays between 48-57
	for i := range srune {
		if int(srune[i]) <= 57 && int(srune[i]) >= 48 {
			result = result*10 + int(srune[i]) - '0'
		} else {
			return 0
		}
	}
	return result
}
