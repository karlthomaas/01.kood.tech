package piscine

func Atoi(s string) int {
	srune := []rune(s)
	result := 0
	// integer unicode stays between 48-57
	// - 45
	// + 43
	negative_state := false
	for i := range srune {
		if i == 0 {
			if int(srune[i]) == 45 {
				negative_state = true
			} else if int(srune[i]) == 43 {
			} else {
				result = result*10 + int(srune[i]) - '0'
			}
		} else if int(srune[i]) <= 57 && int(srune[i]) >= 48 {
			result = result*10 + int(srune[i]) - '0'
		} else {
			return 0
		}
	}
	if negative_state == true {
		return -result
	} else {
		return result
	}
}
