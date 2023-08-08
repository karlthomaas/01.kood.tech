package piscine

func IsPrime(nb int) bool {
	count := 0

	// Kuna tegu on algnumbritega, siis jagu arv AINULT ühe JA endaga(kokku: 2)!!
	// Kui tegu ei ole algnumbriga siis jagub rohkem kui kahega!

	for i := 0; i <= nb; i++ {
		if nb%i == 0 {
			count++
		}
	}
	if count == 2 {
		return true
	} else {
		return false
	}
}
