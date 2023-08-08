package piscine

func ToUpper(s string) string {
	new_str := "" // new string will be stored here

	for _, char := range s {
		// When it's lower ascii
		if char >= 97 && char <= 122 {
			new_str += string(char - 32)
		} else {
			new_str += string(char)
		}
	}
	return new_str
}
