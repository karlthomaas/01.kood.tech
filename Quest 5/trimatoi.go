package piscine

func Atoi(s string) int {
	srune := []rune(s)
	result := 0
	// integer unicode stays between 48-57
	// - 45
	// + 43
	negative_state := false
	integer_presence := false

	for i := range srune {

		// check's the first rune in the list
		// check's if the rune is negative
		if integer_presence == false {
			if int(srune[i]) == 45 {
				negative_state = true
			}
		}

		if int(srune[i]) <= 57 && int(srune[i]) >= 48 {
			integer_presence = true
			result = result*10 + int(srune[i]) - '0'

		} else {
			// do nothing
		}
	}
	if negative_state == true {
		return -result
	} else {
		return result
	}
}

func TrimAtoi(s string) int {
	a := Atoi(s)
	return a
}
