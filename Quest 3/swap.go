package piscine

func Swap(a *int, b *int) {
	temp_a := *a
	temp_b := *b

	*b = temp_a
	*a = temp_b
}
