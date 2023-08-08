package piscine

func ToLower(s string) string {
	new_str := ""
	for _, char := range s {
		if char >= 65 && char <= 90 {
			new_str += string(char + 32)
		} else {
			new_str += string(char)
		}
	}
	return new_str
}
