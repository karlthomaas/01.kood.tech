package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, elem := range tab {
		if f(elem) {
			counter++
		}
	}
	return counter
}
