package piscine

/* Write a function ListAt that takes a pointer to the list l and an int pos as parameters. This function should return the NodeL in the position pos of the linked list l.

In case of error the function should return nil.*/
func ListAt(l *NodeL, nbr int) *NodeL {
	index := 0
	count := 0
	head := l

	for head != nil {
		index++
		head = head.Next
	}

	if nbr <= index {
		for l != nil {
			if count == nbr {
				return l
			}
			count++
			l = l.Next
		}
	}
	return nil
}
