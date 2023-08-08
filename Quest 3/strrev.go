package piscine

func StrRev(s string) string {
	slist := []rune(s) // Rune list, which stores each character
	nstring := ""      // String, where new chars will be added!
	// for loop which will iterate the rune list from the end and
	// append new string with each new character
	for i := len(s) - 1; i >= 0; i-- {
		nstring = nstring + string(slist[i])
	}
	return nstring
}
