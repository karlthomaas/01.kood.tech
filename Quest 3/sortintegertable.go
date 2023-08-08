package piscine

func SortIntegerTable(table []int) []int {
	old_table := table
	new_table := []int{} // Empty integer list
	table_length := len(old_table)
	finished := false
	for len(old_table) > 0 {
		if finished == true {
			break
		}
		// Starts to iterate through the list, takes first value
		// and if the next value is lower/higher then just appends ins
		min_value := old_table[0]
		for index, value := range old_table {
			if value < min_value {
				min_value = value
			}
			// appends
			value_table := []int{min_value}
			new_table = append(value_table, new_table...)
			// removes
			// old_table = remove(old_table, index) // Truncate slice.
			old_table = append(old_table[:index], old_table[index+1:]...)
			if len(new_table) == table_length {
				finished = true
			}
			break
		}
	}
	return new_table
}
