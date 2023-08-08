package piscine

func Map(f func(int) bool, a []int) []bool {
	var bools []bool
	for _, nr := range a {
		boolean := f(nr)
		bools = append(bools, boolean)
	}
	return bools
}
