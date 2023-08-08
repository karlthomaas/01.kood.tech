package piscine

func ForEach(f func(int), a []int) {
	for _, nr := range a {
		f(nr)
	}
}
