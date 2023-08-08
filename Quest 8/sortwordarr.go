package piscine

func SortWordArr(a []string) {
	var newList []string
	var lowest_char string
	var lowest_index int

	tempList := a
	aLength := len(a)

	for {

		if len(newList) == aLength {
			for index, value := range newList {
				a[index] = value
			}
			break
		}

		lowest_char = tempList[0]
		lowest_index = 0

		for index, char := range tempList {
			if char < lowest_char {
				lowest_char = char
				lowest_index = index
			}
		}

		newList = append(newList, string(lowest_char))
		tempList = append(tempList[lowest_index+1:], tempList[:lowest_index]...)
	}
	a = newList
}
