package piscine

func BasicAtoi(s string) int {
	srune := []rune(s)
	result := 0
	for i := range srune {
		result = result*10 + int(srune[i]) - '0'
	}
	return result
}
