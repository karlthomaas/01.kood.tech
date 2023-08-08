package piscine

func NRune(s string, n int) rune {
	a := []rune(s)
	for s, e := range a {
		if s == n-1 {
			return e
		}
	}
	return 0
}
