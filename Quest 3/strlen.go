package piscine

func StrLen(s string) int {
	rune_counter := 0
	for index := range s {
		index++ // place holder, can't have two blanked index || char
		rune_counter++
	}
	return rune_counter
}
