package piscine

func Prime(nb int) bool {
	// if the number is negative or equal to 1 ->
	// Get's added another number (in order to find prime)
	if nb <= 1 {
		return false
	}
	// starts to "for" iterate until the target nb < i*i
	for i := 2; i*i <= nb; i++ {
		// When the number isn't prime number -> Starts to search for another one
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	// Loops until returns
	for {
		if Prime(nb) {
			return nb
		}
		// if returns false, then this section activates ->
		// means that the nb was not prime number ..
		nb++
	}
}
