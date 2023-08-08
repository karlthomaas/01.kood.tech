package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if string(char) >= "a" && string(char) <= "z" || string(char) >= "A" && string(char) <= "Z" || string(char) >= "0" && string(char) <= "9" {
			continue
		} else {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(IsAlpha("Hello! How are you?"))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha("Whatsthis4"))
// }
