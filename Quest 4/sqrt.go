package piscine

func Sqrt(nb int) int {
	for nr := 0; nr <= nb; nr++ {
		if nr*nr == nb {
			return nr
		}
	}
	return 0
}
