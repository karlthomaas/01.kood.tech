package piscine

/*
Write a function ListReverse that reverses the order of the elements of a given linked list l.
*/
func ListReverse(l *List) {
	current := l.Head
	var prev *NodeL
	prev = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
